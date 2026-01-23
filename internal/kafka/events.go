package kafka

import "time"

type BidEvent struct {
	RequestID  string    `json:"request_id"`
	CampaignID int       `json:"campaign_id,omitempty"`
	CreativeID int       `json:"creative_id,omitempty"`
	Price      float64   `json:"price,omitempty"`
	Outcome    string    `json:"outcome"`
	Timestamp  time.Time `json:"ts"`
}

type WinEvent struct {
	CampaignID int       `json:"campaign_id"`
	Price      float64   `json:"price"`
	Timestamp  time.Time `json:"ts"`
}
