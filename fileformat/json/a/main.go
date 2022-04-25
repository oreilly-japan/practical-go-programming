package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	// json-main
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		// ...
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		// ...
	}
	var r ip
	// レスポンスを ip 構造体にデコード
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		// ...
	}
	fmt.Printf("%+v\n", r)
	// json-main
}

// json-struct-ip
type ip struct {
	Origin string `json:"origin"`
}

// json-struct-ip
