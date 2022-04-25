package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type ip struct {
	Origin string `json:"origin"`
	URL    string `json:"url"`
}

// decode-start
func main() {
	s := `{"origin":"255.255.255.255","url":"https://httpbin.org/get"}`
	var resp ip
	if err := json.Unmarshal([]byte(s), &resp); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}

// decode-end
