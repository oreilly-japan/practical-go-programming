package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	yes := 0
	no := 0
	mux := http.NewServeMux()
	mux.Handle("/asset/", http.StripPrefix("/asset/", http.FileServer(http.Dir("."))))
	mux.HandleFunc("/poll/y", func(w http.ResponseWriter, r *http.Request) {
		yes++
	})
	mux.HandleFunc("/poll/n", func(w http.ResponseWriter, r *http.Request) {
		no++
	})
	mux.HandleFunc("/result", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "賛成: %d, 反対: %d", yes, no)
	})
	log.Fatal(http.ListenAndServe(":8080", mux))
}
