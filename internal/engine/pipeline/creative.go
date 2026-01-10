package pipeline

import "gitlab.com/radix-ads/rtb/internal/domain"

func (p *Pipeline) findCreative(winner *domain.BidCandidate) domain.Creative {
	for _, campaign := range p.campaigns {
		if campaign.ID != winner.CampaignID {
			continue
		}

		for _, creative := range campaign.Creatives {
			if creative.ID != winner.CreativeID {
				return creative
			}
		}
	}

	return domain.Creative{}
}
