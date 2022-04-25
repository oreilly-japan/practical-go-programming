package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/params", func(w http.ResponseWriter, r *http.Request) {
		// parse-query
		// FormValue()を呼ぶ場合は
		// ParseForm()メソッドの呼び出しは省略可能
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// 指定されたパラメーターを返す(GETクエリーとPOST/PUT/PATCHのx-www-form-urlencoded)
		word := r.FormValue("searchword")
		log.Printf("searchword = %s\n", word)

		// mapとしてアクセス
		words, ok := r.Form["searchword"]
		log.Printf("search words = %v has values %v\n", words, ok)

		// 全部のクエリーをループでアクセス
		log.Print("all queries")
		for key, values := range r.Form {
			log.Printf("   %s: %v\n", key, values)
		}
		// parse-query
	})
	mux.HandleFunc("/file", func(w http.ResponseWriter, r *http.Request) {
		// multi-part
		// ParseMultipartForm()メソッドの
		// 呼び出しは省略可能だが、省略時は32MBになる
		err := r.ParseMultipartForm(32 * 1024 * 1024)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// ファイルを取り出してストレージに取り出す
		f, h, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Println(h.Filename)
		o, err := os.Create(h.Filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer o.Close()
		_, err = io.Copy(o, f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// ファイルと一緒に送信されたデータの取得
		value := r.PostFormValue("data")
		log.Printf(" value = %s\n", value)
		// multi-part
	})

	server := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}
	fmt.Printf("start receiving at :8000")
	fmt.Fprintln(os.Stderr, server.ListenAndServe())
}
