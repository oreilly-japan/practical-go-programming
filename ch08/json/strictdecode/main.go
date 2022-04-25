package main

import (
	"bytes"
	"encoding/json"
	"os"
)

func main() {
	s, _ := os.ReadFile("rectangle.json")
	// unknown-decode
	var rect Rectangle
	d := json.NewDecoder(bytes.NewReader(s))
	d.DisallowUnknownFields()
	if err := d.Decode(&rect); err != nil {
		// エラーハンドリング
	}
	// unknown-decode
}
