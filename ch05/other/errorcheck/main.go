package main

import "errors"

// error-main
func main() {
	// エラーハンドリングを忘れている
	a()
}

func a() error {
	return errors.New("some error")
}

// error-main
