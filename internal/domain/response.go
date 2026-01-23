package domain

type BidResponse struct {
	RequestID  string
	Price      float64
	Creative   Creative
	CampaignID int
}
