package store

import "gitlab.com/radix-ads/rtb/internal/domain"

type Snapshot struct {
	Campaigns    []domain.Campaign
	CampaignByID map[int]*domain.Campaign
}

func BuildSnapshot(campaigns []domain.Campaign) *Snapshot {
	index := make(map[int]*domain.Campaign, len(campaigns))

	for i := range campaigns {
		campaign := &campaigns[i]
		index[campaign.ID] = campaign
	}

	return &Snapshot{
		Campaigns:    campaigns,
		CampaignByID: index,
	}
}
