package main

import "log"

func main() {
	// two-types
	// ゼロ値
	var empty string

	// 値を設定かつ右辺から型がわかる場合
	one := 1
	// two-types

	// no-use-var
	// 短縮記法の方が良い
	var two = 2
	// no-use-var

	// both-ok
	// varを使っても使わなくてもいい
	var three int8 = 3
	four := int16(4)
	// both-ok

	log.Println(empty, one, two, three, four)
}