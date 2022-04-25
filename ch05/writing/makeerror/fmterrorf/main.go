package main

import (
	"fmt"
	"log"
)

func main() {
	if err := validate(-1); err != nil {
		log.Fatal(err)
	}
}

// errors-error
func validate(length int) error {
	if length <= 0 {
		return fmt.Errorf("length must be greater than 0, lentgh = %d", length)
	}

	// ...
	return nil
}

// errors-error
