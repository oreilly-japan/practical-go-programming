package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("pgx", "host=localhost port=15432 user=testuser dbname=testdb password=pass sslmode=disable")
	if nil != err {
		log.Fatal(err)
	}

	ctx := context.Background()
	value, err := fetchCapacity(ctx, "test")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("value: ", value)
}

// error-continue-start
func fetchCapacity(ctx context.Context, key string) (int, error) {
	var capacity int
	query := `SELECT value FROM parameter_master WHERE key = $1;`
	err := db.QueryRowContext(ctx, query, key).Scan(&capacity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// レコードが存在しなかった場合はデフォルト値を設定し、処理を継続。ログも出力する
			log.Printf("fetch capacity not found, using default capacity, key = %s", key)
			return 10, nil
		}
		return -1, err
	}
	return capacity, nil
}

// error-continue-end
