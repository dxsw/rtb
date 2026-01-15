package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	WinsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "dsp",
			Subsystem: "bidder",
			Name:      "wins_total",
			Help:      "Total wins",
		},
		[]string{"campaign_id"},
	)

	SpendTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "dsp",
			Subsystem: "bidder",
			Name:      "spend_total",
			Help:      "Total spend",
		},
		[]string{"campaign_id"},
	)
)

func RegisterWin() {
	Registry.MustRegister(WinsTotal, SpendTotal)
}
