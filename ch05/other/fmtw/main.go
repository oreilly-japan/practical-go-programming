package main

import (
	"errors"
	"fmt"
)

// error-main
func main() {
	err := errors.New("some error")
	fmt.Printf("%w", err)
}

// error-main
