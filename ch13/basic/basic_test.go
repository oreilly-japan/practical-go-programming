package testsample

import "testing"

func Add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	got := Add(1, 2)
	if got != 3 {
		t.Errorf("expect 3, but %d", got)
	}
}
