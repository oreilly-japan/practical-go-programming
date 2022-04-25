package main

import "fmt"

var (
	a = log("a", c+b)
	b = log("b", f())
	c = log("c", f())
	d = log("d", 3)
)

func f() int {
	d++
	return d
}

func log(name string, i int) int {
	fmt.Println(name)
	return i
}

func main() {
	fmt.Println(a)
}
