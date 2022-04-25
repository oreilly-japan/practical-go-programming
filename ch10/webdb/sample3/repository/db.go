package repository

import (
	"database/sql"
	"fmt"
	"log"
)

// グローバルな変数として宣言して repository パッケージ内で利用
// *sql.DB はゴルーチンセーフなのでグローバル変数として宣言してOK
//
// 非公開の変数としておき、他のパッケージから変更できないようにしておくと安心
var db *sql.DB

func init() {
	var err error

	dbConn := fmt.Sprintf("host=localhost port=15432 user=test dbname=testdb password=pass sslmode=disable")
	db, err = sql.Open("postgres", dbConn)
	if err != nil {
		log.Fatalln("database:", err)
	}
}

type User struct {
	id, name string
}

func FetchUser(id string) (*User, error) {
	// db を使ったデータベース接続の処理
	// a.db.QueryRowContext(...) など
	return &User{ /* ... */ }, nil
}
