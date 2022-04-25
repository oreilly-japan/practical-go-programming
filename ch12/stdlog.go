package main

import (
	"log"
)

func main() {
	log.Println("ログ出力をします")
	// output: 2020/02/13 22:26:27 ログ出力をします
	n := 10
	s := "文字列"
	c := 1 + 2i
	log.Printf("%d, %sなどを使って変数出力もできます。%vはどんな型でも表示します", n, s, c)
}
