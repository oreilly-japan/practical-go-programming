package multiopts

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

// test-start
func TestX(t *testing.T) {
	type X struct {
		NumExport   int
		numUnExport int
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}

	num1 := X{100, -1, time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Now()}
	num2 := X{999, -111, time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), time.Now()}

	opts := []cmp.Option{
		cmpopts.IgnoreUnexported(X{}),
		cmpopts.IgnoreFields(X{}, "CreatedAt", "UpdatedAt"),
	}

	if diff := cmp.Diff(num1, num2, opts...); diff != "" {
		t.Errorf("X value is mismatch (-num1 +num2):%s\n", diff)
	}
}

// test-end
