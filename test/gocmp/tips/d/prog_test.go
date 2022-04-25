package a

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

// test-start
func TestX(t *testing.T) {
	type X struct {
		NumExport int
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	num1 := X{-1, time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Now()}
	num2 := X{-1, time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), time.Now()}

	opt := cmpopts.IgnoreFields(X{}, "CreatedAt", "UpdatedAt")

	if diff := cmp.Diff(num1, num2, opt); diff != "" {
		t.Errorf("X value is mismatch (-num1 +num2):%s\n", diff)
	}
}

// test-end
