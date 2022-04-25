package main

import (
	"encoding/json"
	"fmt"
)

// invalid-json-marshal
type user struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	// 関数型が含まれる場合、エンコードはエラーになる
	X func() `json:"x"`
}

// invalid-json-marshal

func main() {
	u := user{
		UserID:   "001",
		UserName: "gopher",
		X:        func() { fmt.Println("I'm gopher!") },
	}
	b, err := json.Marshal(u)
	if err != nil {
		// エラーハンドリング
		panic(err)
	}
	fmt.Println(string(b))
}
