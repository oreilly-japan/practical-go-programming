package main

import (
	"errors"
	"log"
)

// error-main
func main() {
	_, err := a()
	if err != nil {
		log.Fatal(err)
	}

	b()
	// この err は a() から返ってきたエラーで b() の戻り値のエラーではありません
	if err != nil {
		log.Fatal(err)
	}
}

func a() (string, error) {
	return "", nil
}

func b() error {
	return errors.New("b error")
}

// error-main
