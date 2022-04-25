package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type CreatedColumns struct {
	CreatedAt      time.Time `json:"created_at"`
	CreatedTraceID string    `json:"created_trace_id"`
}

type UpdatedColumns struct {
	CreatedColumns
	UpdatedAt      time.Time `json:"updated_at"`
	UpdatedTraceID string    `json:"updated_trace_id"`
	Revision       int64     `json:"revision"`
}

type Customer struct {
	CustomerID   string `json:"customer_id"`
	CustomerName string `json:"customer_name"`
	UpdatedColumns
}

func main() {

	c := Customer{
		CustomerID:   "0001",
		CustomerName: "オライリー太郎",
		UpdatedColumns: UpdatedColumns{
			CreatedColumns: CreatedColumns{ // ネストで宣言になってしまう
				CreatedAt:      time.Now(),
				CreatedTraceID: "550e8400-e29b-41d4-a716-446655440000",
			},
			UpdatedAt:      time.Now(),
			UpdatedTraceID: "550e8400-e29b-41d4-a716-446655440000",
			Revision:       1,
		},
	}

	b, _ := json.Marshal(c)
	fmt.Printf("%+v", string(b))

}
