package domain

type Campaign struct {
	ID           int
	AdvertiserID int
	Targeting    Targeting
	Budget       Budget
	BidStrategy  BidStrategy
	Creatives    []Creative
}
