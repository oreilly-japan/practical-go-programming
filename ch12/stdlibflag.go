package main

import (
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.SetPrefix("🐙 ")
	log.Println("ログ出力をします")
	// output: 🐙 2020/02/13 22:36:19 sample/flag.go:10: ログ出力をします
}
