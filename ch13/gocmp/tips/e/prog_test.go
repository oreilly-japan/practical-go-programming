package a

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

// test-start
func TestX(t *testing.T) {
	type Item struct {
		HashKey   string
		CreatedAt time.Time
	}
	type X struct {
		Items []Item
	}

	x1 := X{[]Item{
		{HashKey: "aaa", CreatedAt: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)},
		{HashKey: "bbb", CreatedAt: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)},
	}}
	x2 := X{[]Item{
		{HashKey: "aaa", CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)},
		{HashKey: "bbb", CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)},
	}}

	opt := cmpopts.IgnoreFields(Item{}, "CreatedAt")

	if diff := cmp.Diff(x1, x2, opt); diff != "" {
		t.Errorf("X value is mismatch (-x1 +x2):%s\n", diff)
	}
}

// test-end
