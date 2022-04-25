package server

import (
	"fmt"
	"net/http"
	"server/repository"
)

func AppHandler(w http.ResponseWriter, r *http.Request) {
	// repositoryパッケージの関数を用いてデータベースにアクセス
	// 直接 repository.db は参照しない(できない)
	u, err := repository.FetchUser("dummy_id")
	if err != nil {
		// ...
	}
	fmt.Fprintf(w, "handler with database! user=%v", *u)
}
