package main

import "errors"

func a() error {
	return errors.New("some error")
}

// error-main
func main() {
	_ = a()
}

// error-main
