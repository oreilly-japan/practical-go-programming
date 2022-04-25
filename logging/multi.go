package main

import (
	"io"
	"log"
	"os"
)

func main() {
	file, _ := os.Create("log.txt")
	log.SetOutput(io.MultiWriter(file, os.Stderr))
	log.Println("ファイルと標準エラー出力に同時に出力します")
}
