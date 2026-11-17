package clickhouse

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	ch "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
	conn   ch.Conn
}

func NewConsumer(brokers []string, topic, groupID, chDNS string) (*Consumer, error) {
	conn, err := ch.Open(&ch.Options{
		Addr: []string{chDNS},
		Auth: ch.Auth{
			Database: os.Getenv("CLICKHOUSE_DB"),
			Username: os.Getenv("CLICKHOUSE_USER"),
			Password: os.Getenv("CLICKHOUSE_PASSWORD"),
		},
		DialTimeout: 10 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		GroupID:  groupID,
		MinBytes: 1e3,
		MaxBytes: 1e6,
	})

	return &Consumer{reader: reader, conn: conn}, nil
}

func (c *Consumer) Start(ctx context.Context) {
	go func() {
		for {
			m, err := c.reader.ReadMessage(ctx)
			if err != nil {
				if ctx.Err() != nil {
					return
				}
				log.Printf("kafka fetch error: %v", err)
				continue
			}

			var event map[string]any
			if err = json.Unmarshal(m.Value, &event); err != nil {
				log.Printf("json error: %v", err)
				_ = c.reader.CommitMessages(ctx, m)
				continue
			}

			if err = c.insertEvent(ctx, event); err != nil {
				log.Printf("clickhouse insert error: %v", err)
				continue
			}

			_ = c.reader.CommitMessages(ctx, m)
		}
	}()
}

func (c *Consumer) insertEvent(ctx context.Context, event map[string]any) error {
	const query = `INSERT INTO dsp.events (event_type, request_id, impression_id, campaign_id, creative_id, price, outcome,
                                           exchange, site_domain, app_bundle, device_type, country, timestamp)
                   VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	return c.conn.Exec(ctx, query,
		event["event_type"],
		event["request_id"],
		event["imp_id"],
		toUInt32(event["campaign_id"]),
		toUInt32(event["creative_id"]),
		toFloat64(event["price"]),
		event["outcome"],
		event["exchange"],
		event["site_domain"],
		event["app_bundle"],
		event["device_type"],
		event["country"],
		parseTime(event["ts"]))
}

func (c *Consumer) Close() error {
	return c.reader.Close()
}
