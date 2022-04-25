package main

import (
	"fmt"
	"log"
	"time"
)

// basic-struct
type Book struct {
	Title      string
	Author     string
	Publisher  string
	ReleasedAt time.Time
	ISBN       string
}

// basic-struct

func use() {
	// use-struct
	jst, _ := time.LoadLocation("Asia/Tokyo")
	book := Book{
		Title:      "Real World HTTP",
		Author:     "渋川よしき",
		Publisher:  "オライリー・ジャパン",
		ISBN:       "4873119030",
		ReleasedAt: time.Date(2017, time.June, 14, 0, 0, 0, 0, jst),
	}
	fmt.Println(book.Title)
	// use-struct

	b := &Book{
		Title: "Mithril",
	}
	fmt.Println(b.Title)
	fmt.Println((*b).Title)
	b2 := &b
	fmt.Println((**b2).Title)
}

// method
func (b Book) GetAmazonURL() string {
	return "https://amazon.co.jp/dp/" + b.ISBN
}

// method

func nested() {
	// nested-struct-1
	type Book struct {
		Title  string
		Author struct {
			FirstName string
			LastName  string
		}
		Publisher  string
		ReleasedAt time.Time
	}
	// nested-struct-1

	jst, _ := time.LoadLocation("Asia/Tokyo")

	// nested-struct-declare-1
	book := Book{
		Title: "Real World HTTP",
		Author: struct { // インスタンス化するときも定義が必要
			FirstName string
			LastName  string
		}{
			FirstName: "よしき",
			LastName:  "渋川",
		},
		Publisher:  "オライリー・ジャパン",
		ReleasedAt: time.Date(2017, time.June, 14, 0, 0, 0, 0, jst),
	}
	// nested-struct-declare-1

	fmt.Println(book)
}

func nested2() {
	// nested-struct-2
	type Author struct {
		FirstName string
		LastName  string
	}

	type Book struct {
		Title      string
		Author     Author
		Publisher  string
		ReleasedAt time.Time
	}
	// nested-struct-2
}

func anonymous() {
	// anonymous-struct
	jst, _ := time.LoadLocation("Asia/Tokyo")
	book := struct {
		Title      string
		Author     string
		Publisher  string
		ISBN       string
		ReleasedAt time.Time
	}{
		Title:      "Real World HTTP",
		Author:     "渋川よしき",
		Publisher:  "オライリー・ジャパン",
		ISBN:       "4873119030",
		ReleasedAt: time.Date(2017, time.June, 14, 0, 0, 0, 0, jst),
	}
	// anonymous-struct
	log.Println(book)
}

func main() {
	use()
}
