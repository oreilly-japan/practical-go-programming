package main

import (
	"log"
	"net/http"
)

// panic-recovery-start
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// recover を使ってハンドラーで発生した panic から復帰
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// panic-recovery-end

// panic-handler-start
func healthz(w http.ResponseWriter, r *http.Request) {
	// ハンドラー内で意図しない panic が発生...
	panic("panic occur")
}

func main() {
	// panic が発生しても別のリクエストを受け付けることができる
	http.Handle("/healthz", Recovery(http.HandlerFunc(healthz)))
	http.ListenAndServe(":8888", nil)
}

// panic-handler-end
