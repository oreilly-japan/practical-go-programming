package a

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// test-start
func TestX(t *testing.T) {
	type X struct {
		Numbers []int
	}

	num1 := X{[]int{1, 2, 3, 4, 5}}
	num2 := X{[]int{5, 4, 3, 2, 1}}

	if diff := cmp.Diff(num1, num2); diff != "" {
		t.Errorf("X value is mismatch (-num1 +num2):%s\n", diff)
	}
}

// test-end
