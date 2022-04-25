package main

import (
	"fmt"

	"go.uber.org/dig"
)

// init-dig
var Container *dig.Container

type RandomGenerator interface {
	Generate() float64
}

type fakeRandomGenerator struct {}

func (f fakeRandomGenerator) Generate() float64 {
	return 1.0
}

func init () {
	Container = dig.New()

	Container.Provide(func() RandomGenerator {
		return &fakeRandomGenerator{}
	})
}
// init-dig

func main() {
	// exec-dig
	// Invokeを呼び出すと、引数が満たされた状態で関数が実行される
	Container.Invoke(func(rg RandomGenerator) {
		fmt.Println("乱数:", rg.Generate())
	})
	// exec-dig
}