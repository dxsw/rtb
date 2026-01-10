package openrtb

import "gitlab.com/radix-ads/rtb/internal/domain"

func mapImpression(impression Impression) domain.Impression {
	return domain.Impression{
		ID:       impression.ID,
		Banner:   mapBanner(impression.Banner),
		Video:    mapVideo(impression.Video),
		Native:   mapNative(impression.Native),
		BidFloor: impression.BidFloor,
		Secure:   impression.Secure == 1,
	}
}

func mapBanner(banner *Banner) *domain.Banner {
	if banner == nil {
		return nil
	}

	return &domain.Banner{
		Width:  banner.W,
		Height: banner.H,
	}
}

func mapVideo(video *Video) *domain.Video {
	if video == nil {
		return nil
	}

	return &domain.Video{
		Width:       video.W,
		Height:      video.H,
		MIMEs:       video.MIMEs,
		MinDuration: video.MinDur,
		MaxDuration: video.MaxDur,
	}
}

func mapNative(native *Native) *domain.Native {
	if native == nil {
		return nil
	}

	return &domain.Native{
		Version: "1.2",
	}
}

func mapSite(site *Site) *domain.Site {
	if site == nil {
		return nil
	}

	return &domain.Site{
		ID:     site.ID,
		Domain: site.Domain,
		Page:   site.Page,
	}
}

func mapApp(app *App) *domain.App {
	if app == nil {
		return nil
	}

	return &domain.App{
		ID:     app.ID,
		Name:   app.Name,
		Bundle: app.Bundle,
	}
}

func mapDevice(device Device) domain.Device {
	return domain.Device{
		UA:         device.UA,
		IP:         device.IP,
		OS:         device.OS,
		OSVersion:  device.OSVersion,
		Make:       device.Make,
		Model:      device.Model,
		DeviceType: device.DeviceType,
	}
}

func mapUser(user User) domain.User {
	return domain.User{
		ID:       user.ID,
		BuyerUID: user.BuyerUID,
	}
}

func mapRegs(regs Regs) domain.Regs {
	return domain.Regs{
		COPPA: regs.COPPA == 1,
		GDPR:  regs.Ext.GDPR == 1,
	}
}
