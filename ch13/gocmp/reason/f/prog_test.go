package e

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type User struct {
	UserID    string
	UserName  string
	Languages []string
}

// test-start
func TestTom(t *testing.T) {
	tom := User{UserID: "0001", UserName: "Tom", Languages: []string{"Java", "Go"}}
	tom2 := User{UserID: "0001", UserName: "Tom", Languages: []string{"Ruby", "Go"}}

	if diff := cmp.Diff(tom, tom2); diff != "" {
		t.Errorf("User value is mismatch (-tom +tom2):\n%s", diff)
	}
}

// test-end
