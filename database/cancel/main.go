package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	db, err := sql.Open("pgx", "user=testuser password=pass host=localhost port=5432 dbname=testdb sslmode=disable")
	if err != nil {
		panic(err)
	}

	// cancel-start
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// クエリーはコンテキストによって5秒後にキャンセルされます
	if _, err := db.ExecContext(ctx, "SELECT pg_sleep(100)"); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			log.Println("canceling query")
		} else {
			// エラーハンドリング
		}
	}
	// cancel-end

	// 非同期で別ゴルーチンによるクエリーのキャンセルが完了するのを待つ
	time.Sleep(10 * time.Second)

	log.Println("query finish")
}
