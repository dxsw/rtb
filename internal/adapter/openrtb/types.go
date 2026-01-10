package openrtb

type BidResponse struct {
	ID      string    `json:"id"`
	SeatBid []SeatBid `json:"seatbid"`
}

type SeatBid struct {
	Bid []Bid `json:"bid"`
}

type Bid struct {
	ID    string  `json:"id"`
	ImpID string  `json:"impid"`
	Price float64 `json:"price"`

	Adm  string `json:"adm"`
	CrID string `json:"crid"`
}

type BidRequest struct {
	ID     string       `json:"id"`
	Imp    []Impression `json:"imp"`
	Site   *Site        `json:"site"`
	App    *App         `json:"app"`
	Device Device       `json:"device"`
	User   User         `json:"user"`
	Regs   Regs         `json:"regs"`
}

type Impression struct {
	ID       string  `json:"id"`
	Banner   *Banner `json:"banner"`
	Video    *Video  `json:"video"`
	Native   *Native `json:"native"`
	BidFloor float64 `json:"bidfloor"`
	Secure   int     `json:"secure"`
}

type Banner struct {
	W int `json:"w"`
	H int `json:"h"`
}

type Video struct {
	W      int      `json:"w"`
	H      int      `json:"h"`
	MIMEs  []string `json:"mimes"`
	MinDur int      `json:"minduration"`
	MaxDur int      `json:"maxduration"`
}

type Native struct {
	Request string `json:"request"`
}

type Site struct {
	ID     string `json:"id"`
	Domain string `json:"domain"`
	Page   string `json:"page"`
}

type App struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Bundle string `json:"bundle"`
}

type Device struct {
	UA         string `json:"ua"`
	IP         string `json:"ip"`
	OS         string `json:"os"`
	OSVersion  string `json:"osv"`
	Make       string `json:"make"`
	Model      string `json:"model"`
	DeviceType int    `json:"devicetype"`
}

type User struct {
	ID       string `json:"id"`
	BuyerUID string `json:"buyeruid"`
}

type Regs struct {
	COPPA int `json:"coppa"`
	Ext   struct {
		GDPR int `json:"gdpr"`
	} `json:"ext"`
}
