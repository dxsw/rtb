package http

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"gitlab.com/radix-ads/rtb/internal/adapter/openrtb"
	"gitlab.com/radix-ads/rtb/internal/engine"
	"gitlab.com/radix-ads/rtb/internal/fraud"
	"gitlab.com/radix-ads/rtb/internal/infra/kafka"
	"gitlab.com/radix-ads/rtb/internal/metrics"
)

type BidHandler struct {
	Engine   *engine.Engine
	Fraud    *fraud.Guard
	Producer *kafka.Producer
}

func (h *BidHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	request, err := openrtb.ParseRequest(body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	metrics.BidsTotal.Inc()

	isFraud, reason := h.Fraud.IsFraud(request)
	if isFraud {
		metrics.FraudTotal.WithLabelValues(reason).Inc()

		_ = h.Producer.Publish(r.Context(), request.ID, kafka.BidEvent{
			EventType: "bid",
			RequestID: request.ID,
			Outcome:   "fraud",
			Reason:    reason,
			Timestamp: time.Now(),
		})

		w.WriteHeader(http.StatusNoContent)
		return
	}

	start := time.Now()

	response := h.Engine.Process(r.Context(), request)

	metrics.BidLatency.Observe(float64(time.Since(start).Milliseconds()))

	if response == nil {
		metrics.NoBidTotal.Inc()

		_ = h.Producer.Publish(r.Context(), request.ID, kafka.BidEvent{
			EventType:    "bid",
			RequestID:    request.ID,
			ImpressionID: impressionID(request),
			Outcome:      "nobid",
			SiteDomain:   siteDomain(request),
			AppBundle:    appBundle(request),
			Timestamp:    time.Now(),
		})

		w.WriteHeader(http.StatusNoContent)
		return
	}

	_ = h.Producer.Publish(r.Context(), request.ID, kafka.BidEvent{
		EventType:    "bid",
		RequestID:    request.ID,
		ImpressionID: request.Impressions[0].ID,

		CampaignID: response.CampaignID,
		CreativeID: response.Creative.ID,
		Price:      response.Price,
		Outcome:    "win",

		SiteDomain: siteDomain(request),
		AppBundle:  appBundle(request),

		Timestamp: time.Now(),
	})

	openRTBResponse := openrtb.BuildResponse(request, response)
	if openRTBResponse == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	data, err := json.Marshal(openRTBResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
