package server

import (
	"database/sql"
	"net/http"
)

type app struct {
	db *sql.DB
}

func NewApp(db *sql.DB) *app {
	return &app{db: db}
}

func (a *app) AppHandler(w http.ResponseWriter, r *http.Request) {
	// db を使ったデータベース接続の処理
	// a.db.QueryRowContext(...) など
	w.Write([]byte("handler with database!"))
}
