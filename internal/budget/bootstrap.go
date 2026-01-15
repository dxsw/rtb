package budget

import "gitlab.com/radix-ads/rtb/internal/domain"

func BuildStates(campaigns []domain.Campaign) map[int]*State {
	states := make(map[int]*State, len(campaigns))

	for _, campaign := range campaigns {
		states[campaign.ID] = &State{}
	}

	return states
}
