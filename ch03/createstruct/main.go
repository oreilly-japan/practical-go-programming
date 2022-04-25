package main

import (
	"fmt"
)

// Person 構造体
type Person struct {
	FirstName string
	LastName  string
}

// factory-func

// NewPerson ファクトリー関数
func NewPerson(first, last string) *Person {
	return &Person{
		FirstName: first,
		LastName:  last,
	}
}

// factory-func

func main() {
	// create-struct
	// newで作成（あまり使わない）
	p1 := new(Person)

	// var変数宣言で作成
	var p2 Person

	// 複合リテラルで初期化
	p3 := &Person{
		FirstName: "三成",
		LastName:  "石田",
	}
	// create-struct

	// not-create-struct
	// var変数宣言でも、ポインター型の場合は作成されない
	var p4 *Person
	// not-create-struct


	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(p4)

	/*
		// composite-literal
		// こうも書ける
		p5 := &Person{
			"清正", "加藤",
		}
		// が、もしPerson構造体に新しいフィールドが追加されるとエラーに
		// composite-literal
	*/
}
