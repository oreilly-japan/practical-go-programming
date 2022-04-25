package main

import (
	"fmt"

	_ "inittest2/c"
)

func init() {
	fmt.Println("main.init")
}

func main() {
	fmt.Println("main")
}
