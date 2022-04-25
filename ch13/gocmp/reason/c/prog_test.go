package main

import (
	"reflect"
	"testing"
)

type User struct {
	UserID    string
	UserName  string
	Languages []string
}

// test-start
func TestTom(t *testing.T) {
	tom := User{UserID: "0001", UserName: "Tom", Languages: []string{"Java", "Go"}}
	tom2 := User{UserID: "0001", UserName: "Tom", Languages: []string{"Java", "Go"}}

	if !reflect.DeepEqual(tom, tom2) {
		t.Errorf("User tom is mismatch, tom=%v, tom2=%v", tom, tom2)
	}
}

// test-end
