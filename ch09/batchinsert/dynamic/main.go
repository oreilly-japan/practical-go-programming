package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	ctx := context.Background()
	users := []struct {
		userID, userName string
	}{
		{userID: "0001", userName: "Gopher"},
		{userID: "0002", userName: "Ferris"},
		{userID: "0003", userName: "Duke"},
	}

	db, err := sql.Open("pgx", "host=localhost port=5432 user=testuser dbname=testdb password=pass sslmode=disable")
	if nil != err {
		panic(err)
	}
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	// batch-dynamic-start
	// VALUES に記述する文字列を組み立てる
	valueStrings := make([]string, 0, len(users))
	valueArgs := make([]interface{}, 0, len(users)*2 /* 追加したい項目数 (user_id, user_name) */)

	number := 1
	for _, u := range users {
		valueStrings = append(valueStrings, fmt.Sprintf(" ($%d, $%d)", number, number+1))
		valueArgs = append(valueArgs, u.userID)
		valueArgs = append(valueArgs, u.userName)
		number += 2
	}

	query := fmt.Sprintf("INSERT INTO users (user_id, user_name) VALUES %s;", strings.Join(valueStrings, ","))
	if _, err := db.ExecContext(ctx, query, valueArgs...); err != nil {
		// エラーハンドリング
	}
	// batch-dynamic-end

	if err := tx.Commit(); err != nil {
		// エラーハンドリング
	}
}
