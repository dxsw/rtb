package store

import (
	"sync/atomic"
	"time"
)

type AtomicSnapshot struct {
	value atomic.Value
}

func NewAtomicSnapshot(initial *Snapshot) *AtomicSnapshot {
	atomicSnapshot := &AtomicSnapshot{}
	atomicSnapshot.value.Store(initial)
	return atomicSnapshot
}

func (s *AtomicSnapshot) Load() *Snapshot {
	return s.value.Load().(*Snapshot)
}

func (s *AtomicSnapshot) Store(snapshot *Snapshot) {
	s.value.Store(snapshot)
}

func StartReloadLoop(atomic *AtomicSnapshot, loader Loader, interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for range ticker.C {
			campaigns, err := loader.LoadCampaigns()
			if err != nil {
				// TODO: log error
				continue
			}

			snapshot := BuildSnapshot(campaigns)
			atomic.Store(snapshot)
		}
	}()
}
