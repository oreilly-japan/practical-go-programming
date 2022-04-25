package main

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

var (
	burst   = 10
	limiter = NewIpRateLimiter(rate.Every(time.Second/1), burst)
)

func main() {
	http.Handle("/test", limit(http.HandlerFunc(hello)))
	http.ListenAndServe(":18888", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world!")
}

func limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limiter := limiter.Get(r.RemoteAddr)
		if !limiter.Allow() {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
