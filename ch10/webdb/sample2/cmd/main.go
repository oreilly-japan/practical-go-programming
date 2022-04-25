package main

import (
	"database/sql"
	"log"
	"net/http"
	"server"
)

func main() {
	dbConn := "host=localhost port=15432 user=test dbname=testdb password=pass sslmode=disable"
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		log.Fatalln("database:", err)
	}

	// mainパッケージ内で必要な構造体を生成
	app := server.NewApp(db)

	// func(http.ResponseWriter, *http.Request)を実装した
	// 構造体のメソッドを呼び出す
	http.HandleFunc("/test", app.AppHandler)

	log.Fatal(http.ListenAndServe(":18888", nil))
}
