package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// decode-start
type ip struct {
	Origin string `json:"origin"`
	// url と非公開の変数とするとマッピングできない
	url string `json:"url"`
}

func main() {
	s := `{"origin":"255.255.255.255","url":"https://httpbin.org/get"}`
	var resp ip
	if err := json.Unmarshal([]byte(s), &resp); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
	// {Origin:255.255.255.255 url:}
}

// decode-end
