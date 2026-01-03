package domain

type Video struct {
	MIMEs       []string
	MinDuration int
	MaxDuration int

	Width  int
	Height int

	StartDelay int
	Protocols  []int
	Linearity  int
	Skippable  *bool
}
