package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// systemcolumns-struct-start
type CreatedColumns struct {
	CreatedAt      time.Time `json:"created_at"`
	CreatedTraceID string    `json:"created_trace_id"`
}

type UpdatedColumns struct {
	UpdatedAt      time.Time `json:"updated_at"`
	UpdatedTraceID string    `json:"updated_trace_id"`
	Revision       int64     `json:"revision"`
}

// systemcolumns-struct-end

// systemcolumns-customer-start
type Customer struct {
	CustomerID   string `json:"customer_id"`
	CustomerName string `json:"customer_name"`
	CreatedColumns
	UpdatedColumns
}

// systemcolumns-customer-end

func main() {

	// systemcolumns-main-start
	c := Customer{
		CustomerID:   "0001",
		CustomerName: "オライリー太郎",
		CreatedColumns: CreatedColumns{
			CreatedAt:      time.Now(),
			CreatedTraceID: "550e8400-e29b-41d4-a716-446655440000",
		},
		UpdatedColumns: UpdatedColumns{
			UpdatedAt:      time.Now(),
			UpdatedTraceID: "550e8400-e29b-41d4-a716-446655440000",
			Revision:       1,
		},
	}

	b, _ := json.Marshal(c)
	fmt.Printf("%+v", string(b))
	// systemcolumns-main-end

}
