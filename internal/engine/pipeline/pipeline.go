package pipeline

import (
	"gitlab.com/radix-ads/rtb/internal/domain"
	"gitlab.com/radix-ads/rtb/internal/engine/matcher"
	"gitlab.com/radix-ads/rtb/internal/engine/selection"
)

type Pipeline struct {
	campaigns []domain.Campaign
}

func New(campaigns []domain.Campaign) *Pipeline {
	return &Pipeline{
		campaigns: campaigns,
	}
}

func (p *Pipeline) Process(request *domain.BidRequest) *domain.BidResponse {
	candidates := matcher.Match(request, p.campaigns)
	if len(candidates) == 0 {
		return nil
	}

	winner := selection.SelectWinner(candidates)
	if winner == nil {
		return nil
	}

	return &domain.BidResponse{
		RequestID: request.ID,
		Price:     winner.Price,
		Creative:  p.findCreative(winner),
	}
}
