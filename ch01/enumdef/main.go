package main

import (
	"fmt"
)

// go-generate
//go:generate stringer -type=CarOption
//go:generate stringer -type=CarType
// go-generate

// enumer-generate
//go:generate enumer -type=CarOption -json
//go:generate enumer -type=CarType -json
// enumer-generate

// typedef1
type CarType int

// typedef1

// constant
const (
	Sedan CarType = iota + 1
	Hatchback
	MPV
	SUV
	Crossover
	Coupe
	Convertible
)

// constant

// typedef2
type CarOption uint64

const (
	GPS CarOption = 1 << iota
	AWD
	SunRoof
	HeatedSeat
	DriverAssist
)

// typedef2

// iota-behavior
const (
	a = iota // 0
	b        // 1
	c        // 2
	_        // 3だが使われない
	// 空行は無視
	d        // 4
	e = iota // 5
)

const (
	f = iota // 再び0
	g        // 1
	h        // 2
)

// iota-behavior

func main() {
	// use-enum1
	// 値を設定
	var t CarType
	t = SUV
	// use-enum1

	// use-enum2
	// 組み合わせて設定
	var o CarOption
	o = SunRoof | HeatedSeat

	if o&SunRoof != 0 {
		fmt.Println("サンルーフ付き")
	}
	// use-enum2

	fmt.Println(a, b, c, d, e, f, g, h, t, o)

	// go-generate-test
	c := Convertible
	fmt.Printf("握力王の愛車は%sです\n", c)
	// 握力王の愛車はConvertibleです
	// go-generate-test
}
