package main

import (
	"encoding/json"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	h := Hello{Message: "Hello world!"}
	enc := json.NewEncoder(w)
	enc.Encode(&h)
}

var Handler = http.HandlerFunc(Handle)

type Hello struct {
	Message string `json:"message"`
}
