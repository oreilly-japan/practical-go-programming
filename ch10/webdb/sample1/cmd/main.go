package main

import (
	"database/sql"
	"log"
	"net/http"
	"server"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type withDBFunc func(*sql.DB, http.ResponseWriter, *http.Request)

// withDBはfunc(w http.ResponseWriter, r *http.Request)の関数に
// データベースハンドラーを扱えるようにしたラッパーです。
//
// http.HandlerFuncに直接渡せるシグネチャは
// 	func(w http.ResponseWriter, r *http.Request)
// であるため、*sql.DB はクロージャで参照します。
func withDB(db *sql.DB, f withDBFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(db, w, r)
	}
}

func main() {
	dbConn := "host=localhost port=15432 user=test dbname=testdb password=pass sslmode=disable"
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		log.Fatalln("database:", err)
	}
	http.HandleFunc("/test", withDB(db, server.AppHandler))

	log.Fatal(http.ListenAndServe(":18888", nil))
}
