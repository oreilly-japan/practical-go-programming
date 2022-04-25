package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// json-main
	u := user{
		UserID:   "001",
		UserName: "gopher",
	}
	b, _ := json.Marshal(u)
	fmt.Println(string(b))
	// json-main
}
