package main

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	http.Handle("/test", LimitHandler(http.HandlerFunc(hello)))
	http.ListenAndServe(":18888", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world!")
}

// limit-start
// 1 秒ごとに 1 回までに制限するリミッター、バーストは 10
var limiter = rate.NewLimiter(rate.Every(time.Second/1), 10)

func LimitHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// limit-end
