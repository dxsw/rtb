package matcher

import "gitlab.com/radix-ads/rtb/internal/domain"

func Match(request *domain.BidRequest, campaigns []domain.Campaign) []domain.BidCandidate {
	if len(request.Impressions) == 0 {
		return nil
	}

	impression := request.Impressions[0]

	candidates := make([]domain.BidCandidate, 0, 8)

	for _, campaign := range campaigns {
		for _, creative := range campaign.Creatives {

			if !creativeMatchesImpression(&creative, &impression) {
				continue
			}

			price := campaign.BidStrategy.BaseCPM
			if price < impression.BidFloor {
				continue
			}

			candidates = append(candidates, domain.BidCandidate{
				CampaignID: campaign.ID,
				CreativeID: creative.ID,
				Price:      price,
			})
		}
	}

	return candidates
}
