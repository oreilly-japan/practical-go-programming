package main

import (
	"fmt"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world!")
	}
	http.HandleFunc("/test", handler)

	// デフォルトのマルチプレクサに追加
	http.Handle("/asset/", http.StripPrefix("/asset/", http.FileServer(http.Dir("asset"))))

	http.ListenAndServe(":18888", nil)
}
