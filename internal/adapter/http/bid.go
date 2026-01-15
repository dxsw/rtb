package http

import (
	"encoding/json"
	"io"
	"net/http"

	"gitlab.com/radix-ads/rtb/internal/adapter/openrtb"
	"gitlab.com/radix-ads/rtb/internal/engine"
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

	response := h.Engine.Process(request)
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
