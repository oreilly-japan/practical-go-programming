package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// sample-struct
type Record struct {
	ProcessID string `json:"process_id"`
	DeletedAt JSTime `json:"deleted_at"`
}

type JSTime time.Time

func (t JSTime) MarshalJSON() ([]byte, error) {
	tt := time.Time(t)
	if tt.IsZero() {
		return []byte("null"), nil
	}
	v := strconv.Itoa(int(tt.UnixMilli()))
	return []byte(v), nil
}

// sample-struct

func main() {
	// execute
	r := &Record{
		ProcessID: "0001",
		DeletedAt: JSTime(time.Now()),
	}

	b, _ := json.Marshal(r)
	fmt.Println(string(b))
	// {"process_id":"0001","delete_ttl":1639450000919}
	// execute
}
