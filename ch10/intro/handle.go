package main

import (
	"net/http"
)

// hello-handler-start
func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}

// hello-handler-end

func main() {
	// serve-mux-start
	// HTTPサーバーのパス /hello に Hello 関数の
	// アプリケーションハンドラーを紐付けている
	http.HandleFunc("/hello", Hello)
	// serve-mux-end

	// listen-start

	http.ListenAndServe(":18888", nil)
	// listen-end
}
