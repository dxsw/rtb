package budget

import (
	"sync/atomic"

	"gitlab.com/radix-ads/rtb/internal/store"
)

type Checker struct {
	snapshot *store.Snapshot
	states   map[int]*State
}

func (c *Checker) CanBid(campaignID int, price float64) bool {
	campaign := c.snapshot.CampaignByID[campaignID]
	if campaign == nil {
		return false
	}

	state := c.states[campaign.ID]
	if state == nil {
		return true
	}

	cost := uint64(price * 100)

	if campaign.Budget.DailyCap > 0 {
		if atomic.LoadUint64(&state.DailySpent)+cost > uint64(campaign.Budget.DailyCap*100) {
			return false
		}
	}

	if campaign.Budget.TotalCap > 0 {
		if atomic.LoadUint64(&state.TotalSpent)+cost > uint64(campaign.Budget.TotalCap*100) {
			return false
		}
	}

	return true
}
