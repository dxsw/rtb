package domain

type Impression struct {
	ID       string
	Banner   *Banner
	Video    *Video
	Native   *Native
	BidFloor float64
	Secure   bool
}
