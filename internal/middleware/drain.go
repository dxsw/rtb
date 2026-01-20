package middleware

import (
	"net/http"
	"sync"
	"sync/atomic"
)

type Drain struct {
	wg     sync.WaitGroup
	closed atomic.Bool
}

func (d *Drain) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if d.closed.Load() {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}

		d.wg.Add(1)
		defer d.wg.Done()

		next.ServeHTTP(w, r)
	})
}

func (d *Drain) Close() {
	d.closed.Store(true)
}

func (d *Drain) Wait() {
	d.wg.Wait()
}
