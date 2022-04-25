package main

import (
	"fmt"
)

// def-embed-struct
type Book struct {
	Title string
	ISBN  string
}

func (b Book) GetAmazonURL() string {
	return "https://amazon.co.jp/dp/" + b.ISBN
}

type OreillyBook struct {
	Book
	ISBN13 string
}

func (o OreillyBook) GetOreillyURL() string {
	return "https://www.oreilly.co.jp/books/" + o.ISBN13 + "/"
}

// def-embed-struct

func main() {
	// use-embed-struct
	ob := OreillyBook{
		ISBN13: "9784873119038",
		Book: Book{
			Title: "Real World HTTP",
		},
	}
	// Bookのメソッドが利用可能
	fmt.Println(ob.GetAmazonURL())
	// OreillyBookのメソッドも利用可能
	fmt.Println(ob.GetOreillyURL())
	// use-embed-struct
}

/*
// def-compat
type OreillyBook struct {
	Book   Book   // 型名と同じ名前のフィールドがあるのとほぼ同じ
	ISBN13 string
}

// こうも書ける
fmt.Println(ob.Book.GetAmazonURL())
// def-compat
*/

/*
func test() {
	// multiple-embed
	type Database struct {
		Address string
	}
	type FileStorage struct {
		Address string
	}
	type WebServiceConfig struct {
		Database
		FileStorage
	}
	ws := WebServiceConfig{
		Database: Database{
			Address: "Databaseのフィールド",
		},
		FileStorage: FileStorage{
			Address: "FileStorageのフィールド",
		},
	}
	// ambiguous selector ws.Addressというエラー
	fmt.Println(ws.Address)
	// このように埋め込んだ構造体のどちらを参照したいか明示する
	fmt.Println(ws.Database.Address)
	// multiple-embed
}
*/
