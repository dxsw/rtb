package pipeline

import (
	"context"

	"gitlab.com/radix-ads/rtb/internal/budget"
	"gitlab.com/radix-ads/rtb/internal/domain"
	"gitlab.com/radix-ads/rtb/internal/engine/matcher"
	"gitlab.com/radix-ads/rtb/internal/engine/selection"
	"gitlab.com/radix-ads/rtb/internal/pacing"
	"gitlab.com/radix-ads/rtb/internal/store"
)

type Pipeline struct {
	campaigns []domain.Campaign
}

func New(campaigns []domain.Campaign) *Pipeline {
	return &Pipeline{
		campaigns: campaigns,
	}
}

func (p *Pipeline) Process(ctx context.Context, request *domain.BidRequest, snapshot *store.Snapshot, checker *budget.Checker, pacer *pacing.Pacer) *domain.BidResponse {
	select {
	case <-ctx.Done():
		return nil
	default:
	}

	candidates := matcher.Match(request, snapshot.Campaigns)
	if len(candidates) == 0 {
		return nil
	}

	filtered := candidates[:0]
	for _, candidate := range candidates {
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		if checker.CanBid(candidate.CampaignID, candidate.Price) && pacer.ShouldBid(candidate.CampaignID, candidate.Price) {
			filtered = append(filtered, candidate)
		}
	}

	winner := selection.SelectWinner(filtered)
	if winner == nil {
		return nil
	}

	return &domain.BidResponse{
		RequestID:  request.ID,
		Price:      winner.Price,
		Creative:   p.findCreative(winner),
		CampaignID: winner.CampaignID,
	}
}
