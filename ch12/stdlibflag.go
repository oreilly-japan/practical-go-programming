package main

import (
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.SetPrefix("π ")
	log.Println("γ­γ°εΊεγγγΎγ")
	// output: π 2020/02/13 22:36:19 sample/flag.go:10: γ­γ°εΊεγγγΎγ
}
