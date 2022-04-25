package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

// def-func
func calc(x, y int) int {
	return x + y
}

// def-func

// def-func2
func calcAge(y int, m time.Month, d int) (age int, err error) {
	b := time.Date(y, m, d, 0, 0, 0, 0, time.Local)
	n := time.Now()
	if b.After(n) {
		err = errors.New("誕生日が未来です")
		return
	}
	for {
		b = time.Date(y+age+1, m, d, 0, 0, 0, 0, time.Local)
		if b.After(n) {
			return
		}
		age++
	}
}

// def-func2

// def-func4
func doCalc(x, y int, f func(int, int) int) {
	fmt.Println()
}

// def-func4

func main() {
	fmt.Println(calcAge(1980, time.November, 14))
	// def-func3
	m := func(x, y int) int {
		return x * y
	}

	// def-func5
	// 名前付きで定義した関数以外に、無名関数として定義したものを渡せる
	doCalc(10, 20, m)
	// その場で定義した関数も渡せる
	doCalc(10, 10, func(x, y int) int {
		return x * y
	})
	// def-func5

	// def-func3
	fmt.Println(m(1, 2))

	// def-func6
	// ファイルを開く
	f, err := os.Create("sample.txt")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	// この関数のスコープを抜けたら自動でファイルをクローズ
	defer f.Close()
	io.WriteString(f, "hello world")
	// def-func6

}
