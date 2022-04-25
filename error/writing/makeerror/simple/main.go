package main

import (
	"errors"
)

// errors-error
var ErrNotFound = errors.New("not found")

func findBook(isbn string) (*Book, error) {
	// ...
	// 値が取得できなかったため ErrNotFound を返す
	return nil, ErrNotFound
}

// errors-error

type Book struct {
}
