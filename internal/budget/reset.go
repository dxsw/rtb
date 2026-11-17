package budget

import (
	"context"
	"sync/atomic"
	"time"
)

type DailyResetter struct {
	States map[int]*State
}

func (r *DailyResetter) Start(ctx context.Context) {
	go func() {
		for {
			duration := timeUntilNextDay()

			timer := time.NewTimer(duration)
			select {
			case <-timer.C:
				r.reset()
			case <-ctx.Done():
				timer.Stop()
				return
			}
		}
	}()
}

func (r *DailyResetter) reset() {
	for _, state := range r.States {
		atomic.StoreUint64(&state.DailySpent, 0)
	}
}

func timeUntilNextDay() time.Duration {
	now := time.Now()
	next := time.Date(
		now.Year(),
		now.Month(),
		now.Day()+1,
		0, 0, 0, 0,
		now.Location(),
	)
	return time.Until(next)
}
