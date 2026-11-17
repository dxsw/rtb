package budget

import "sync/atomic"

type Service struct {
	states map[int]*State
}

func NewService(states map[int]*State) *Service {
	return &Service{states: states}
}

func (s *Service) OnWin(campaignID int, price float64) {
	cost := uint64(price * 100)

	state := s.states[campaignID]
	if state == nil {
		return
	}

	atomic.AddUint64(&state.DailySpent, cost)
	atomic.AddUint64(&state.TotalSpent, cost)
}
