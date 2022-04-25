package main

import (
	"fmt"

	_ "inittest/c"
	_ "inittest/b"
)

func init() {
	fmt.Println("main.init")
}

func main() {
	fmt.Println("main")
}
