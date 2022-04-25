package collection

import "time"

func Expires(d time.Time) {
	time.Now().AddDate(0, 1, 0)
}
