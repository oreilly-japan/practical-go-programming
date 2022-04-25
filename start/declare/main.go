package main

import (
	"fmt"
	"strconv"
)

/*
// literal
1     // 数字
1.23  // 浮動小数点数
"abc" // 文字列
`今日は
いい天気` // バッククォートも文字列（改行可能）
true  // ブール値
nil   // nil: 無効な参照先を表す値
// literal
*/

func main() {
	// variable
	// 整数型の変数
	var num1 int = 123

	// 右辺で型が決まるなら型名は不要
	var num2 = 123

	// variable

	// short-variable
	// 変数宣言と代入を同時に行う短縮記法
	// ただし関数内部でのみ利用可能
	num3 := 123
	// short-variable

	// convert
	var i int = 123
	// 数値同志の変換は、かっこでくくり型を前置します
	var f float64 = float64(i)
	// 64ビットOSで64ビットのintと、int64も明示的な変換が必要です
	var i64 int64 = int64(i)
	// boolへの変換は比較演算子を使います
	var b bool = i != 0
	// convert

	// strconv-convert
	// 文字列との変換はstrconvパッケージを利用
	in := 12345
	// strconvの数値入力はint64, uint64, float64なので
	// それ以外の変数を使う時は型変換が必要
	s := strconv.FormatInt(int64(in), 10)
	// s = "12345"

	// Parse系はエラー変換失敗時にエラーを返す
	// 成功時のerrはnil
	f, err := strconv.ParseFloat("12.3456", 64)
	// f = 12.3456
	// err = nil

	// strconv-convert

	fmt.Println(num1, num2, num3, i, f, i64, b, s, f, err)
	pointer()
}

func pointer() {
	// use-pointer
	// ポインターの参照先となる普通の変数
	var i int = 10

	// ポインターを格納する変数（デフォルト値はnil）
	var p *int

	// pにはiのアドレスが入る
	p = &i

	// p経由でiの値を取得
	fmt.Println(*p)
	// use-pointer
}
