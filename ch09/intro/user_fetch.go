package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var db *sql.DB

func main() {
	dbConn, err := sql.Open("pgx", "host=localhost port=5432 user=testuser dbname=testdb password=pass sslmode=disable")
	if nil != err {
		log.Fatal(err)
	}
	db = dbConn

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err)
	}

	dumpALLUser(ctx)

}

func dumpALLUser(ctx context.Context) {
	// fetch-rows-start
	type User struct {
		UserID    string
		UserName  string
		CreatedAt time.Time
	}

	rows, err := db.QueryContext(ctx, `SELECT user_id, user_name, created_at FROM users ORDER BY user_id;`)
	if err != nil {
		log.Fatalf("query all users: %v", err)
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		var (
			userID, userName string
			createdAt        time.Time
		)
		if err := rows.Scan(&userID, &userName, &createdAt); err != nil {
			log.Fatalf("scan the user: %v", err)
		}
		users = append(users, &User{
			UserID:    userID,
			UserName:  userName,
			CreatedAt: createdAt,
		})
	}
	if err := rows.Close(); err != nil {
		log.Fatalf("rows close: %v", err)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("scan users: %v", err)
	}
	// fetch-rows-end

	fmt.Println(users)
}

// FetchUser は users テーブルからユーザIDに紐づくユーザを取得
func FetchUser(ctx context.Context, userID string) {
	type User struct {
		UserID    string
		UserName  string
		CreatedAt time.Time
	}

	// fetch-row-start
	var (
		userName  string
		createdAt time.Time
	)
	row := db.QueryRowContext(ctx, `SELECT user_name, created_at FROM users WHERE user_id = $1;`, userID)
	err := row.Scan(&userName, &createdAt)
	if err != nil {
		log.Fatalf("query row(user_id=%s): %v", userID, err)
	}
	u := User{
		UserID:    userID,
		UserName:  userName,
		CreatedAt: createdAt,
	}
	// fetch-row-end
	fmt.Println(u)
}
