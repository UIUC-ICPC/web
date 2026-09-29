package middleware

import (
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// TODO: add tests to verify rate limit works
func RateLimit(next http.Handler) http.Handler {
	var mu sync.Mutex
	counts := make(map[string]int)
	reset := time.Now().Add(time.Minute)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, "Invalid client address", http.StatusBadRequest)
			return
		}

		mu.Lock()
		now := time.Now()
		if !now.Before(reset) {
			counts = make(map[string]int)
			reset = now.Add(time.Minute)
		}

		allowed := counts[ip] < 60
		if allowed {
			counts[ip]++
		}
		retryAfter := int(reset.Sub(now).Seconds()) + 1
		mu.Unlock()

		if !allowed {
			w.Header().Set("Retry-After", strconv.Itoa(retryAfter))
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
