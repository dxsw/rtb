package http

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"gitlab.com/radix-ads/rtb/internal/adapter/openrtb"
	"gitlab.com/radix-ads/rtb/internal/engine"
	"gitlab.com/radix-ads/rtb/internal/metrics"
)

type BidHandler struct {
	Engine *engine.Engine
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

	start := time.Now()

	response := h.Engine.Process(request)

	metrics.BidLatency.Observe(float64(time.Since(start).Milliseconds()))

	if response == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

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
