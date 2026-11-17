package domain

type User struct {
	ID       string
	BuyerUID string
	Segments []int
	FreqCaps map[int]int
}
