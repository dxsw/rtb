package middleware

import (
	"context"
	"net/http"
	"time"
)

type Timeout struct {
	Duration time.Duration
}

func (t *Timeout) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), t.Duration)
		defer cancel()

		r = r.WithContext(ctx)

		done := make(chan struct{})
		go func() {
			next.ServeHTTP(w, r)
			close(done)
		}()

		select {
		case <-done:
			return
		case <-ctx.Done():
			w.WriteHeader(http.StatusNoContent)
			return
		}
	})
}
