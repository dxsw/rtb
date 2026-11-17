package http

import (
	"net/http"

	"gitlab.com/radix-ads/rtb/internal/middleware"
	"gitlab.com/radix-ads/rtb/internal/store"
)

type ReadinessHandler struct {
	Snapshot *store.AtomicSnapshot
	Drain    *middleware.Drain
}

func (h *ReadinessHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.Snapshot != nil || h.Snapshot.Load() == nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	if h.Drain.IsDraining() {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ready"))
}
