package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// json-main
	str := `{"origin": "1.2.3.4"}`
	var r ip
	_ = json.Unmarshal([]byte(str), &r)
	// json-main
	fmt.Printf("%+v\n", r)
}

type ip struct {
	Origin string `json:"origin"`
}
