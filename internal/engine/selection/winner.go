package selection

import "gitlab.com/radix-ads/rtb/internal/domain"

func SelectWinner(candidates []domain.BidCandidate) *domain.BidCandidate {
	if len(candidates) == 0 {
		return nil
	}

	winner := candidates[0]

	for i := 1; i < len(candidates); i++ {
		candidate := candidates[i]
		if candidate.Price > winner.Price {
			winner = candidate
		}
	}

	return &winner
}
