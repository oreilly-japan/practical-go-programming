package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// null-encode-start
	type user struct {
		UserID    string   `json:"user_id"`
		UserName  string   `json:"user_name"`
		Languages []string `json:"languages"`
	}

	u := user{
		UserID:   "001",
		UserName: "gopher",
	}
	b, _ := json.Marshal(u)
	fmt.Println(string(b))
	// {"user_id":"001","user_name":"gopher","languages":null}
	// null-encode-end
}
