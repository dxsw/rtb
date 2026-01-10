package openrtb

import (
	"encoding/json"
	"errors"

	"gitlab.com/radix-ads/rtb/internal/domain"
)

func ParseRequest(data []byte) (*domain.BidRequest, error) {
	var openRTBRequest BidRequest
	if err := json.Unmarshal(data, &openRTBRequest); err != nil {
		return nil, err
	}

	if openRTBRequest.ID == "" || len(openRTBRequest.Imp) == 0 {
		return nil, errors.New("invalid OpenRTB request")
	}

	request := &domain.BidRequest{
		ID:          openRTBRequest.ID,
		Impressions: make([]domain.Impression, 0, len(openRTBRequest.Imp)),
		Device:      mapDevice(openRTBRequest.Device),
		User:        mapUser(openRTBRequest.User),
		Site:        mapSite(openRTBRequest.Site),
		App:         mapApp(openRTBRequest.App),
		Regs:        mapRegs(openRTBRequest.Regs),
	}

	for _, impression := range openRTBRequest.Imp {
		request.Impressions = append(request.Impressions, mapImpression(impression))
	}

	return request, nil
}
