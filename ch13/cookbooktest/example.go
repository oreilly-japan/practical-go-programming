package cookbooktest

import "fmt"

// in-start

func F() {
	fmt.Println("example_f")
}

type T struct{}

func (t T) M() {
	fmt.Println("example_m")
}

func (t T) String() string {
	return "example_t"
}

// in-end
