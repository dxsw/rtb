package http

import (
	"encoding/json"
	"net/http"

	"gitlab.com/radix-ads/rtb/internal/budget"
)

type WinNotifyHandler struct {
	Budget *budget.Service
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

	var req winRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.CampaignID == 0 || req.Price <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h.Budget.OnWin(req.CampaignID, req.Price)

	w.WriteHeader(http.StatusNoContent)
}
