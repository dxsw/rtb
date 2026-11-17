package openrtb

import (
	"strconv"

	"gitlab.com/radix-ads/rtb/internal/domain"
)

func BuildResponse(request *domain.BidRequest, response *domain.BidResponse) *BidResponse {
	if response == nil {
		return nil
	}

	bid := Bid{
		ID:    strconv.Itoa(response.Creative.ID),
		ImpID: request.Impressions[0].ID,
		Price: response.Price,
		Adm:   response.Creative.Adm,
		CrID:  strconv.Itoa(response.Creative.ID),
	}

	return &BidResponse{
		ID: request.ID,
		SeatBid: []SeatBid{
			{
				Bid: []Bid{bid},
			},
		},
	}
}
