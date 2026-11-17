package http

import "gitlab.com/radix-ads/rtb/internal/domain"

func siteDomain(req *domain.BidRequest) string {
	if req.Site != nil {
		return req.Site.Domain
	}
	return ""
}

func appBundle(req *domain.BidRequest) string {
	if req.App != nil {
		return req.App.Bundle
	}
	return ""
}

func impressionID(req *domain.BidRequest) string {
	if len(req.Impressions) > 0 {
		return req.Impressions[0].ID
	}
	return ""
}
