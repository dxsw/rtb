package store

import "gitlab.com/radix-ads/rtb/internal/domain"

type StaticLoader struct{}

func NewStaticLoader() *StaticLoader {
	return &StaticLoader{}
}

func (l *StaticLoader) LoadCampaigns() ([]domain.Campaign, error) {
	return []domain.Campaign{
		{
			ID:           1,
			AdvertiserID: 10,
			Budget: domain.Budget{
				DailyCap: 100.0,
			},
			Creatives: []domain.Creative{
				{
					ID:     1,
					Width:  300,
					Height: 250,
					MIME:   "text/html",
					Adm:    "<div>test</div>",
				},
			},
		},
	}, nil
}
