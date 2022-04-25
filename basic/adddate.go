package main

import (
	"fmt"
	"time"
)

func addDate() {
	// next-month-normal
	jst, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Date(2021, 6, 8, 20, 56, 00, 000, jst)
	nextMonth := now.AddDate(0, 1, 0)
	fmt.Println(nextMonth)
	// 2021-07-08 20:56:00 +0900 JST
	// next-month-normal

	// normalization
	normal := time.Date(2021, 6, 31, 00, 00, 00, 000, jst)
	fmt.Println(normal)
	// 2021-07-01 00:00:00 +0900 JST
	// normalization
}

func next() {
	jst, _ := time.LoadLocation("Asia/Tokyo")

	// normalization2
	now := time.Date(2021, 5, 31, 00, 00, 00, 000, jst)
	nextMonth := now.AddDate(0, 1, 0)
	fmt.Println(nextMonth)
	// 2021-07-01 00:00:00 +0900 JST
	// normalization2
}

// calc-next-month
func NextMonth(t time.Time) time.Time {
	year1, month2, day := t.Date()
	first := time.Date(year1, month2, 1, 0, 0, 0, 0, time.UTC)
	year2, month2, _ := first.AddDate(0, 1, 0).Date()
	nextMonthTime := time.Date(year2, month2, day, 0, 0, 0, 0, time.UTC)
	if month2 != nextMonthTime.Month() {
		return first.AddDate(0, 2, -1) // 翌月末
	}
	return nextMonthTime
}

// calc-next-month

func main() {
	addDate()
	next()
}
