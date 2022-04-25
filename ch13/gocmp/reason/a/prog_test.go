package main

import "testing"

// test-start
type User struct {
	UserID   string
	UserName string
}

func TestTom(t *testing.T) {
	tom := User{UserID: "0001", UserName: "Tom"}
	tom2 := User{UserID: "0001", UserName: "Tom"}

	if tom != tom2 {
		t.Errorf("User tom is mismatch, tom=%v, tom2=%v", tom, tom2)
	}
}

// test-end
