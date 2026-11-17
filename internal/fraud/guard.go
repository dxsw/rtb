package fraud

import (
	"net"
	"strings"
	"time"

	"gitlab.com/radix-ads/rtb/internal/domain"
)

type Guard struct {
	DatacenterCIDRs []*net.IPNet
	BlockedDomains  map[string]struct{}
	BlockedBundles  map[string]struct{}
	MaxIPQPS        int
	ipCounters      map[string]*counter
}

type counter struct {
	count int
	reset time.Time
}

func NewGuard() *Guard {
	return &Guard{
		DatacenterCIDRs: []*net.IPNet{},
		BlockedDomains:  map[string]struct{}{},
		BlockedBundles:  map[string]struct{}{},
		MaxIPQPS:        50,
		ipCounters:      make(map[string]*counter),
	}
}

func (g *Guard) IsFraud(req *domain.BidRequest) (bool, string) {
	ip := req.Device.IP
	if ip != "" {
		if g.isOverIPLimit(ip) {
			return true, "ip_rate_limit"
		}
	}

	if ip != "" && g.isDatacenterIP(ip) {
		return true, "datacenter_ip"
	}

	if req.Site != nil {
		siteDomain := strings.ToLower(req.Site.Domain)
		if _, ok := g.BlockedDomains[siteDomain]; ok {
			return true, "blocked_domain"
		}
	}

	if req.App != nil {
		bundle := strings.ToLower(req.App.Bundle)
		if _, ok := g.BlockedBundles[bundle]; ok {
			return true, "blocked_bundle"
		}
	}

	if req.Device.IP == "" || req.Device.UA == "" || req.Impressions == nil || len(req.Impressions) == 0 {
		return true, "invalid_request"
	}

	return false, ""
}

func (g *Guard) isOverIPLimit(ip string) bool {
	now := time.Now()
	c, ok := g.ipCounters[ip]
	if !ok || now.After(c.reset) {
		g.ipCounters[ip] = &counter{count: 1, reset: now.Add(time.Second)}
		return false
	}
	c.count++
	return c.count > g.MaxIPQPS
}

func (g *Guard) isDatacenterIP(ip string) bool {
	parsed := net.ParseIP(ip)
	if parsed == nil {
		return true
	}

	for _, cidr := range g.DatacenterCIDRs {
		if cidr.Contains(parsed) {
			return true
		}
	}

	return false
}
