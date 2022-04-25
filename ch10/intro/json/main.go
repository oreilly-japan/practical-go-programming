package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Comment struct {
	Message  string
	UserName string
}

func main() {
	var mutex = &sync.RWMutex{}
	comments := make([]Comment, 0, 100)

	http.HandleFunc("/comments", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case http.MethodGet: // GETメソッドの処理
			mutex.RLock() // 読み取り時に書き込みがあることを考慮しロックする。本来はDBから読み取る処理を代替
			if err := json.NewEncoder(w).Encode(comments); err != nil {
				http.Error(w, fmt.Sprintf(`{"status":"%s"}`, err), http.StatusInternalServerError)
				return
			}
			mutex.RUnlock()

		case http.MethodPost: // POSTメソッドの処理
			var c Comment
			if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
				http.Error(w, fmt.Sprintf(`{"status":"%s"}`, err), http.StatusInternalServerError)
				return
			}
			mutex.Lock() // 同時に複数アクセスを防ぐためにロック。本来はDBに保存する処理を代替
			comments = append(comments, c)
			mutex.Unlock()

			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"status": "created"}`))

		default: // GET, POST以外は許容しない
			http.Error(w, `{"status":"permits only GET or POST"}`, http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8888", nil)
}
