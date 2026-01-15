package budget

import (
	"sync/atomic"
	"time"
)

type DailyResetter struct {
	States map[int]*State
}

func (r *DailyResetter) Start() {
	go func() {
		for {
			sleepUntilNextDay()
			r.reset()
		}
	}()
}

func (r *DailyResetter) reset() {
	for _, state := range r.States {
		atomic.StoreUint64(&state.DailySpent, 0)
	}
}

func sleepUntilNextDay() {
	now := time.Now()
	next := time.Date(
		now.Year(),
		now.Month(),
		now.Day()+1,
		0, 0, 0, 0,
		now.Location(),
	)
	time.Sleep(time.Until(next))
}
