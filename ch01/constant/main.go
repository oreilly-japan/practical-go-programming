package main

import (
	"database/sql"
	"errors"
	"log"
)

func literal() {
	// literal-sample
	var a = 1       // イコールの右側の1が定数
	var b = 1 + 100 // 定数の中で演算子を使った計算も可能
	// literal-sample
	log.Println(a, b)
}

func bindType() {
	// bind-type
	var a = 1             // aの型は定数のデフォルトの型（ここではint）になる
	var b int32 = 1 + 100 // 明示的に書くことも可能
	function(1000)        // 関数の引数に渡すリテラルも定数だが、関数の引数の型に合わせて型が決定
	// bind-type
	log.Println(a, b)
}

func function(i int) {

}

// non-type-const
const (
	a = 1              // イコールの右側の1が定数
	b = 1 + 2          // 演算もできる
	c = 9223372036854775807 + 1 // uint64を超える数値も定義可能
	d = "hello world"  // 整数以外に、浮動小数点数や文字列やブール型も扱える
	e = iota + 10      // const宣言でのみiotaも使える
)

// non-type-const

// typed-const
type ErrorCode int

const (
	f    int       = 10
	code ErrorCode = 10
)

// typed-const

/*
// typed-const-error
const (
	g int32          = 4294967295 + 1  // これは型の範囲を超えるためエラーになる
	h []int          = []int{1, 2, 3}  // 配列やスライスはエラー
	i map[string]int = map[string]int{ // マップもエラー
		"Tokyo":    10,
		"Kanagawa": 11,
		"Chiba":    12,
	}
	j = function() // 関数の返り値もエラー
)

// typed-const-error
*/

// constant-instance
type errDatabase int

func (e errDatabase) Error() string {
	return "Database Error"
}

const (
	ErrDatabase errDatabase = 0
)

// constant-instance

/*
// err-constant-instance
const (
	// New()関数の返り値は実行時まで決まらないのでconstにできない
	ErrDatabase = errors.New("Database Error")
)
// err-constant-instance
*/

func OpenDB(param string) error {
	return ErrDatabase
}

func main() {
	// use-constant-instance
	err := OpenDB("postgres://localhost:5432")
	if err == ErrDatabase {
		log.Fatal("DB接続エラー")
	}
	// use-constant-instance
	// override-error-instance
	sql.ErrConnDone = errors.New("エラーを乗っ取りもできてしまう")
	// override-error-instance
}
