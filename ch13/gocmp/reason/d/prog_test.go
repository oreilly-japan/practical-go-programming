package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type User struct {
	UserID    string
	UserName  string
	Languages []string
}

func TestTom(t *testing.T) {
	tom := User{UserID: "0001", UserName: "Tom", Languages: []string{"Java", "Go"}}
	tom2 := User{UserID: "0001", UserName: "Tom", Languages: []string{"Java", "Go"}}

	// test-start
	// go-cmp の Diff を使って値を比較
	if diff := cmp.Diff(tom, tom2); diff != "" {
		t.Errorf("User value is mismatch (-tom +tom2):\n%s", diff)
	}
	// test-end
}
