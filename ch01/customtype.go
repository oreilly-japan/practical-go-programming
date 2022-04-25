package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// timestamp-type-def
type JSTimeStamp time.Time

func (t JSTimeStamp) MarshalJSON() ([]byte, error) {
	ot := time.Time(t)
	v := strconv.FormatInt(ot.UnixMilli(), 10)
	return []byte(v), nil
}

// timestamp-type-def

func customType() {
	// timestamp-type-use
	t := JSTimeStamp(time.Now())
	j, _ := json.Marshal(&t)
	fmt.Println(string(j))
	// 1600257607675

	// % node
	// Welcome to Node.js v14.10.1.
	// Type ".help" for more information.
	// > new Date(1600257607675)
	// 2020-09-16T12:00:07.675Z
	// timestamp-type-use
}

func main() {
	customType()
}
