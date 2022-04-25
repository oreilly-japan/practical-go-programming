package main

import (
	"bufio"
	"os"
)

func main() {
	// error-intro
	f, err := os.Open("important.txt")
	if err != nil {
		// エラーハンドリング
	}

	r := bufio.NewReader(f)
	l, err := r.ReadString('\n')
	if err != nil {
		// エラーハンドリング
	}
	// ここではエラーは発生していない

	// error-intro
	_ = l
}
