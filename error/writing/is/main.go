package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

var db sql.DB

func main() {
	handle()
}

func handle() error {
	ctx := context.Background()

	sessionID := "abc"

	// error-handle-start
	isNew := false
	u, err := fetchSession(ctx, sessionID)
	if err != nil {
		// データベースレイヤーから返却されたエラーの詳細を調べてハンドリングできます
		if !errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("fail to fetch session: %w", err)
		}
		isNew = true
	}
	// ...

	// error-handle-end

	_ = u
	_ = isNew

	return nil
}

// error-wrap-start
func fetchSession(ctx context.Context, sessionID string) (*Session, error) {
	var s *Session

	query := `SELECT * FROM sessions WHERE session_id = $1;`
	err := db.QueryRowContext(ctx, query, sessionID).Scan(&s)
	if err != nil {
		return nil, fmt.Errorf("fetch session, sessionID = %s: %w", sessionID, err)
	}

	return s, nil
}

// error-wrap-end

type Session struct {
}
