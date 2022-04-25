package main

import (
	"net/http"
)

func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	http.Handle("/healthz", http.HandlerFunc(Healthz))
	http.ListenAndServe(":8888", nil)
}
