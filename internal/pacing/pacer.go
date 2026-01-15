package pacing

import (
	"math"
	"sync/atomic"
	"time"

	"gitlab.com/radix-ads/rtb/internal/budget"
	"gitlab.com/radix-ads/rtb/internal/store"
)

type Pacer struct {
	Snapshot *store.Snapshot
	States   map[int]*budget.State
	Clock    Clock

	MaxBidFraction float64
	Tolerance      float64
}

func (p *Pacer) ShouldBid(campaignID int, price float64) bool {
	campaign := p.Snapshot.CampaignByID[campaignID]
	if campaign == nil {
		return false
	}

	dailyCap := campaign.Budget.DailyCap
	if dailyCap <= 0 {
		return true
	}

	state := p.States[campaignID]
	if state == nil {
		return true
	}

	spent := float64(atomic.LoadUint64(&state.DailySpent)) / 100.0

	now := p.Clock.Now()

	expected := expectedSpend(dailyCap, now)
	allowed := expected * p.Tolerance

	if spent+price > allowed {
		return false
	}

	maxBid := dailyCap * p.MaxBidFraction
	if price > maxBid {
		return false
	}

	return true
}

func expectedSpend(dailyCap float64, now time.Time) float64 {
	start := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		0, 0, 0, 0,
		now.Location(),
	)

	elapsed := now.Sub(start).Seconds()
	day := 24 * 60 * 60.0

	progress := math.Min(elapsed/day, 1.0)

	return dailyCap * progress
}
