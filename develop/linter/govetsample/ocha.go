package govetsample

import "fmt"

func Drink() {
	// vet-sample-start
	// Printf のフォーマット引数を忘れている
	fmt.Printf("おーい %s")
	// vet-sample-end
}
