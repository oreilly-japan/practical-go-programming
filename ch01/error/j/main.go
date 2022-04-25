package main

import "os"

func main() {
	// file-open
	f, err := os.Open("important.txt")
	if err != nil {
		// Open時に発生したエラーを処理する場合
	}
	defer f.Close()
	// 変数 f を使ったなにかの処理
	// file-open
}
