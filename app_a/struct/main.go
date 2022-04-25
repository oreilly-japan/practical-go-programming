// json-tag
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type Book struct {
	Title      string    `json:"title"`
	Author     string    `json:"author"`
	Publisher  string    `json:"publisher"`
	ReleasedAt time.Time `json:"release_at"`
	ISBN       string    `json:"isbn"`
}

func main() {
	f, err := os.Open("book.json")
	if err != nil {
		log.Fatal("file open error: ", err)
	}
	d := json.NewDecoder(f)
	var b Book
	d.Decode(&b)
	fmt.Println(b)
}

// json-tag

/*
// def-struct
type Book struct {
	Title      string
	Author     string
	Publisher  string
	ReleasedAt time.Time
	ISBN       string
}

// def-struct
*/

func useStruct() {
	// use-struct
	// インスタンス作成（フィールドはすべてゼロ値に初期化）
	var b Book

	// フィールドを初期化しながらインスタンス作成
	b2 := Book{
		Title: "Twisted Network Programming Essentials",
	}

	// フィールドを初期化しながらインスタンス作成
	// 変数にはポインターを格納
	b3 := &Book{
		Title: "カンフーマック――猛獣を飼いならす310の技",
	}

	// use-struct
}
