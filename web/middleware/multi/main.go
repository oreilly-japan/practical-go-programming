package main

import (
	"log"
	"net/http"
)

func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// ハンドラーを呼び出す前後でログを出力するミドルウェア
func MiddlewareLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("start %s\n", r.URL)
		next.ServeHTTP(w, r)
		log.Printf("finish %s\n", r.URL)
	})
}

func main() {
	// middleware-timeout-start
	h := MiddlewareLogging(http.HandlerFunc(Healthz))
	http.Handle("/healthz", http.TimeoutHandler(h, 5, "request timeout"))
	http.ListenAndServe(":8888", nil)
	// middleware-timeout-end
}
