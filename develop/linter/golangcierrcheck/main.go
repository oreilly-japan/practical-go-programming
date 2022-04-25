package main

import "errors"

func main() {
	Something()
}

func Something() error {
	return errors.New("error occurred")
}
