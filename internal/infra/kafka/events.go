package kafka

import "time"

type BidEvent struct {
	EventType    string    `json:"event_type"`
	RequestID    string    `json:"request_id"`
	ImpressionID string    `json:"impression_id,omitempty"`
	CampaignID   int       `json:"campaign_id,omitempty"`
	CreativeID   int       `json:"creative_id,omitempty"`
	Price        float64   `json:"price,omitempty"`
	Outcome      string    `json:"outcome"`
	Reason       string    `json:"reason"` // TODO: save reason in ClickHouse
	Exchange     string    `json:"exchange,omitempty"`
	SiteDomain   string    `json:"site_domain,omitempty"`
	AppBundle    string    `json:"app_bundle,omitempty"`
	DeviceType   string    `json:"device_type,omitempty"`
	Country      string    `json:"country,omitempty"`
	Timestamp    time.Time `json:"timestamp"`
}

type WinEvent struct {
	EventType    string    `json:"event_type"`
	ImpressionID string    `json:"impression_id,omitempty"`
	CampaignID   int       `json:"campaign_id"`
	Price        float64   `json:"price"`
	Exchange     string    `json:"exchange,omitempty"`
	Timestamp    time.Time `json:"timestamp"`
}
