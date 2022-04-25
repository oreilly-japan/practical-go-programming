package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type User struct {
	UserID   string
	UserName string
}

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=testuser dbname=testdb password=pass sslmode=disable")
	if nil != err {
		log.Fatal(err)
	}
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	ctx := context.Background()

	// prepared-start
	users := []User{
		{"0001", "Gopher"},
		{"0002", "Ferris"},
		{"0003", "Duke"},
	}
	// プリペアードステートメントの構築
	stmt, err := tx.PrepareContext(ctx, "INSERT INTO users(user_id, user_name, created_at) VALUES( $1, $2, $3 );")
	if err != nil {
		// エラーハンドリング
	}
	defer stmt.Close()

	for _, u := range users {
		// 構築したプリペアードステートメントにパラメーターをセットし、実行
		if _, err := stmt.ExecContext(ctx, u.UserID, u.UserName); err != nil {
			// エラーハンドリング
		}
	}
	if err := tx.Commit(); err != nil {
		// エラーハンドリング
	}
	// prepared-end
}
