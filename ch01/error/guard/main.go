package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// error-intro
	f, err := os.Open("important.txt")
	if err == nil {
		r := bufio.NewReader(f)
		l, err := r.ReadString('\n')
		if err == nil {
			// 変数 l を使った何らかの処理
			// ここではエラーは発生していない
			fmt.Println(l)
		}
	}
	// エラーハンドリング

	// error-intro
}
