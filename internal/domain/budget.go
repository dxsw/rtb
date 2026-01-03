package domain

type Budget struct {
	DailyLimit float64
	SpentToday float64
}

type PacingState struct {
	Tokens float64
}
