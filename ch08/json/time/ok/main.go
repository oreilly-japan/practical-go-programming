package main

import (
	"encoding/json"
	"log"
	"strconv"
	"time"
)

const format = time.RFC3339

type Hoge struct {
	T ISO8601Time
}

type ISO8601Time time.Time

func (t ISO8601Time) MarshalJSON() ([]byte, error) {
	s := time.Time(t).Format(strconv.Quote(format))
	return []byte(s), nil
}

func main() {
	loc, _ := time.LoadLocation("Asia/Tokyo")

	h := Hoge{T: ISO8601Time(time.Date(2019, time.August, 1, 10, 0, 0, 0, loc))}
	b, err := json.Marshal(&h)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(string(b))
}
