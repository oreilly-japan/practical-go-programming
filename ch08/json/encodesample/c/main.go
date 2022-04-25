package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type user struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
}

func main() {
	// json-main
	var b bytes.Buffer
	u := user{
		UserID:   "001",
		UserName: "gopher",
	}
	_ = json.NewEncoder(&b).Encode(u)
	fmt.Printf("%v\n", b.String())
	// {"user_id":"001","user_name":"gopher"}

	// json-main
}
