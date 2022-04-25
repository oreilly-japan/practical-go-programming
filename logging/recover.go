package main

import (
	"fmt"
)

func main() {
	// go-recover
	defer func() {
		err := recover()
		fmt.Println("in-recover:", err)
	}()

	panic("panic message")
	// in-recover: panic message
	// go-recover
}
