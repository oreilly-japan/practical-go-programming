package main

import "fmt"

// override
type Parent struct{}

func (p Parent) m1() {
	// 親は子に埋め込まれているが、親は子のことは知らないので、親のm2
	// が呼ばれる
	p.m2()
}

func (p Parent) m2() {
	fmt.Println("Parent")
}

type Child struct {
	Parent
}

func (c Child) m2() {
	fmt.Println("Child")
}

// override

func ngOverride() {
	// use-sample
	c := Child{}
	c.m1() // Parent: オーバーライド失敗
	c.m2() // Child
	// use-sample
}

func main() {
	ngOverride()
}
