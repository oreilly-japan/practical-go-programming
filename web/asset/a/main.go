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
	http.Handle("/", http.FileServer(http.Dir(".")))

	http.ListenAndServe(":18888", nil)
}
