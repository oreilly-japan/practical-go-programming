package b

import (
	"fmt"

	_ "inittest/a"
)

func init() {
	fmt.Println("b.init")
}
