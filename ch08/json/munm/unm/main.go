package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// sample-struct
type Record struct {
	ProcessID string `json:"process_id"`
	DeletedAt JSTime `json:"deleted_at"`
}

type JSTime time.Time

// method-start
func (t *JSTime) UnmarshalJSON(data []byte) error {
	var jsonNumber json.Number
	err := json.Unmarshal(data, &jsonNumber)
	if err != nil {
		return err
	}
	unix, err := jsonNumber.Int64()
	if err != nil {
		return err
	}

	*t = JSTime(time.Unix(0, unix))
	return nil
}

// method-end

// sample-struct

func main() {
	// execute
	s := `{"process_id":"001","deleted_at":1639485878560}`
	var r *Record
	if err := json.Unmarshal([]byte(s), &r); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", time.Time(r.DeletedAt).Format(time.RFC3339Nano))
	// execute
}
