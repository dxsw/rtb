package store

import "gitlab.com/radix-ads/rtb/internal/domain"

type Loader interface {
	LoadCampaigns() ([]domain.Campaign, error)
}
