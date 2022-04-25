package main

import "fmt"

func main() {

	condition := true

	// shadow-start
	x := 1
	if condition {
		// コードブロックの内側で変数 x を再宣言できる
		x := 2
		fmt.Println("x =", x)
	}

	// conditionがtrueの場合でも、ここの位置では必ず1になる
	fmt.Println("x =", x)
	// shadow-end
}
