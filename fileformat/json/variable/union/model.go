package union

// union-start
type Response struct {
	Type      string `json:"type"`
	Timestamp int    `json:"timestamp"`
	Payload   `json:"payload"`
}

type Payload struct {
	ID string `json:"id"`

	// "type" が message のときに使うフィールド
	UserID    string  `json:"user_id"`
	Message   string  `json:"message"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`

	// "type" が sensor のときに使うフィールド
	DeviceID  string `json:"device_id"`
	Result    string `json:"result"`
	ProductID string `json:"product_id"`
}

// union-end
