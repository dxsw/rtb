package metrics

import "github.com/prometheus/client_golang/prometheus"

var FraudTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: "dsp",
		Subsystem: "bidder",
		Name:      "fraud_total",
		Help:      "Total fraud filtered bids",
	},
	[]string{"reason"},
)

func RegisterFraud() {
	Registry.MustRegister(FraudTotal)
}
