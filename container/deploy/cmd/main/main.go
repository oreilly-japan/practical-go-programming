package main

import (
	"log"
	"net/http"
	"sampleapp"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", sampleapp.Handler))
}
