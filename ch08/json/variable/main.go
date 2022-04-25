package main

import (
	"encoding/json"
	"log"
	"os"
)

// raw-struct-start
type Response struct {
	Type      string `json:"type"`
	Timestamp int    `json:"timestamp"`
	// Payload を具体的な構造体に展開せず json.RawMessage として保持
	Payload json.RawMessage `json:"payload"`
}

// message のペイロードを扱う構造体
type Message struct {
	ID        string  `json:"id"`
	UserID    string  `json:"user_id"`
	Message   string  `json:"message"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// sensor のペイロードを扱う構造体
type Sensor struct {
	ID        string `json:"id"`
	DeviceID  string `json:"device_id"`
	Result    string `json:"result"`
	ProductID string `json:"product_id"`
}

// raw-struct-end

func main() {

	jsonStr, err := os.ReadFile("message.json")
	if err != nil {
		log.Fatal(err)
	}

	// raw-message-start
	// 一度の json.RawMessage のフィールドとしてデコードし
	// Type の値によってもう一度デコードする
	var r Response
	_ = json.Unmarshal(jsonStr, &r)

	switch r.Type {
	case "message":
		var m Message
		_ = json.Unmarshal(r.Payload, &m)
	case "sensor":
		// ...
	default:
		// ...
	}
	// raw-message-end
}
