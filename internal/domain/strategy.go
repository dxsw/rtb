package domain

type BidStrategyType int

const (
	BidStrategyFixedCPM BidStrategyType = iota
	BidStrategyDynamicCPM
)

type Multipliers struct {
	PCTR float64
	PCVR float64
	Time float64
}

type BidStrategy struct {
	Type BidStrategyType

	BaseCPM float64

	MaxCPM float64
	MinCPM float64

	Multipliers Multipliers
}
