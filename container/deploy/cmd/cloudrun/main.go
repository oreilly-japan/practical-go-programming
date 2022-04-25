package main

import (
	"log"
	"net/http"
	"os"
	"sampleapp"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Fatal(http.ListenAndServe(":"+port, sampleapp.Handler))
}
