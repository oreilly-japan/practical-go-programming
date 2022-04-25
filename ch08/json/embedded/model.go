package main

import "time"

// created-columns-start
type CreatedColumns struct {
	CreatedAt      time.Time `json:"created_at"`
	CreatedTraceID string    `json:"created_trace_id,omitempty"`
}

// created-columns-end
