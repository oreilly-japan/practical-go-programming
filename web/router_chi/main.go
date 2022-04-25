package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	yes := 0
	no := 0
	r := chi.NewRouter()
	r.Post("/poll/{answer}", func(w http.ResponseWriter, r *http.Request) {
		if chi.URLParam(r, "answer") == "y" {
			yes++
		} else {
			no++
		}
	})
	r.Get("/result", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "賛成: %d, 反対: %d", yes, no)
	})
	r.Handle("/asset/*", http.StripPrefix("/asset/", http.FileServer(http.Dir("."))))

	log.Fatal(http.ListenAndServe(":8080", r))
}
