package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
)

// defmethod
// メソッドを追加する型
// 構造体でなくてもプリミティブ型に対する型も可能
type Struct struct {
	v int
}

// レシーバーを持つ関数としてメソッドを定義
func (s Struct) PrintStatus() {
	log.Println("Struct:", s.v)
}

// defmethod

/*
// receiver
// インスタンスのレシーバー
// 値を変更してもインスタンスのフィールドは変更されない
func (s Struct) SetValue(v int) {
	s.v = v
}

// ポインターのレシーバー
// 値を変更できる
func (s *Struct) SetValue(v int) {
	s.v = v
}
// receiver
*/

// struct-with-pointer
// イディオムに反する実装例
type StructWithPointer struct {
	v *int
}

// このメソッドはインスタンスレシーバーだが変更できてしまう
func (a StructWithPointer) Modify() {
	(*a.v) = 10
}

// struct-with-pointer

type Value struct {
	Id   string
	Name string
}

// func-compat
type Handler struct {
	db *sqlx.DB
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	var v Value
	err := h.db.Get(&v, "SELECT * FROM person WHERE id=$1", r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(v)
}

// func-compat
/*
// use-func
func register(h *Handler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/value", h.Get)
	return mux
}
// use-func
*/

// closure
func register(db *sqlx.DB) http.Handler {
	mux := http.NewServeMux()

	// 関数定義は関数内関数で実装
	// 以前レシーバー経由で使っていたデータベースのコネクションは
	// クロージャがキャプチャした関数定義外のregister()の引数を利用
	mux.HandleFunc("/value", func(w http.ResponseWriter, r *http.Request) {
		var v Value
		err := db.Get(&v, "SELECT * FROM person WHERE id=$1", r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(v)
	})
	return mux
}

// closure

func main() {
	// use-struct-with-pointer
	s := StructWithPointer{}
	i := 1
	s.v = &i
	s.Modify()
	fmt.Println(*s.v)
	// 結果は10となって変更されているが、紛らわしいので
	// このようなコードを書いてはいけません
	// use-struct-with-pointer
}
