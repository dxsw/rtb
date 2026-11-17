package matcher

import "gitlab.com/radix-ads/rtb/internal/domain"

func creativeMatchesImpression(creative *domain.Creative, impression *domain.Impression) bool {
	switch creative.Format {

	case domain.CreativeBanner:
		if impression.Banner == nil {
			return false
		}
		return sizeMatch(creative.Width, creative.Height, impression.Banner.Width, impression.Banner.Height)

	case domain.CreativeVideo:
		if impression.Video == nil {
			return false
		}
		return sizeMatch(creative.Width, creative.Height, impression.Video.Width, impression.Video.Height)

	case domain.CreativeNative:
		return impression.Native != nil

	default:
		return false
	}
}
