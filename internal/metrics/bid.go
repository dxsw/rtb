package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	BidsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "dsp",
		Subsystem: "bidder",
		Name:      "bids_total",
		Help:      "Total bid requests",
	})

	NoBidTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "dsp",
		Subsystem: "bidder",
		Name:      "no_bid_total",
		Help:      "Total no-bid responses",
	})

	BidLatency = prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: "dsp",
		Subsystem: "bidder",
		Name:      "bid_latency_ms",
		Help:      "Bid request latency",
		Buckets:   []float64{1, 2, 5, 10, 20, 50},
	})
)

func RegisterBid() {
	Registry.MustRegister(
		BidsTotal,
		NoBidTotal,
		BidLatency,
	)
}
