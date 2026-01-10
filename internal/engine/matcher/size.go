package matcher

func sizeMatch(campaignWidth, campaignHeight, impressionWidth, impressionHeight int) bool {
	return campaignWidth == impressionWidth && campaignHeight == impressionHeight
}
