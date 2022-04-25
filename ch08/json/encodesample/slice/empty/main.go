package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	UserID    string   `json:"user_id"`
	UserName  string   `json:"user_name"`
	Languages []string `json:"languages"`
}

func main() {
	// empty-slice
	u := user{
		UserID:   "001",
		UserName: "gopher",
		// 明示的に空スライスをセット
		Languages: []string{},
	}
	b, _ := json.Marshal(u)
	fmt.Println(string(b))
	// {"user_id":"001","user_name":"gopher","languages":[]}
	// empty-slice
}
