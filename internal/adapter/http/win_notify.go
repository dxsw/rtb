package http

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"gitlab.com/radix-ads/rtb/internal/budget"
	"gitlab.com/radix-ads/rtb/internal/kafka"
	"gitlab.com/radix-ads/rtb/internal/metrics"
)

type WinNotifyHandler struct {
	Budget   *budget.Service
	Producer *kafka.Producer
	Secret   string
}

type winRequest struct {
	CampaignID int     `json:"campaign_id"`
	Price      float64 `json:"price"`
}

func (h *WinNotifyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	signature := r.Header.Get("X-Signature")
	if signature == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !verifyHMAC(h.Secret, body, signature) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var req winRequest
	if err = json.Unmarshal(body, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.CampaignID == 0 || req.Price <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h.Budget.OnWin(req.CampaignID, req.Price)

	_ = h.Producer.Publish(r.Context(), "win", kafka.WinEvent{
		CampaignID: req.CampaignID,
		Price:      req.Price,
		Timestamp:  time.Now(),
	})

	metrics.WinsTotal.WithLabelValues(strconv.Itoa(req.CampaignID)).Inc()
	metrics.SpendTotal.WithLabelValues(strconv.Itoa(req.CampaignID)).Add(req.Price)

	w.WriteHeader(http.StatusNoContent)
}
