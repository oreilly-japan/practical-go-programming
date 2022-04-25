package main

import (
	"log"
)

type any = interface{}

func empty() {
	// use-empty-interface
	// interface{}型はAny型
	var v any = "なんでも入る変数"
	// 整数も入る
	v = 1
	// 浮動小数点数も入る
	v = 3.141592
	// use-empty-interface

	// use-empty-composition
	// interface{}型のスライス
	var slices = []any{
		"関ヶ原",
		1600,
	}
	// interface{}型のmap
	var ieyasu = map[string]any{
		"名前": "徳川家康",
		"生まれ": 1543,
	}
	// use-empty-composition

	log.Println(v, slices, ieyasu)
}

func main() {
	empty()}

