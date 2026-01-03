package domain

type BidRequest struct {
	ID          string
	Impressions []Impression
	Device      Device
	User        User
	Site        *Site
	App         *App
	Regs        Regs
}
