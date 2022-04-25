package server

import (
	"database/sql"
	"net/http"
)

func AppHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// db を使ったデータベース接続の処理
	// db.QueryRowContext(...) など
	w.Write([]byte("handler with database!"))
}
