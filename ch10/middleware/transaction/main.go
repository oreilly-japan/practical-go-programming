package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code // ステータスコードの保存
	lrw.ResponseWriter.WriteHeader(code)
}

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// recover を使ってハンドラで発生した panic から復帰
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// go-middleware-transaction-start
func NewMiddlewareTx(db *sql.DB) func(http.Handler) http.Handler {
	return func(wrappedHandler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tx, _ := db.Begin()
			lrw := NewLoggingResponseWriter(w)
			r = r.WithContext(context.WithValue(r.Context(), "tx", tx))

			wrappedHandler.ServeHTTP(lrw, r)

			statusCode := lrw.statusCode
			if 200 <= statusCode && statusCode < 400 {
				log.Println("transaction committed")
				tx.Commit()
			} else {
				log.Print("transaction rolling back due to status code: ", statusCode)
				tx.Rollback()
			}
		})
	}
}

// go-middleware-transaction-end

// go-middleware-transaction-extract-start
func extractTx(r *http.Request) *sql.Tx {
	tx, ok := r.Context().Value("tx").(*sql.Tx)
	if !ok {
		panic("transaction middleware is not supported")
	}
	return tx
}

// go-middleware-transaction-extract-end

// go-middleware-transaction-main-start
func main() {
	db := openMyDB()          // *sql.DBを取得
	tx := NewMiddlewareTx(db) // ミドルウェアを生成

	http.Handle("/comments", tx(Recovery(http.HandlerFunc(Comments))))
	http.ListenAndServe(":8888", nil)
}

func Comments(w http.ResponseWriter, r *http.Request) {
	tx := extractTx(r)
	// DBアクセス処理
}

// go-middleware-transaction-main-end

func openMyDB() *sql.DB {
	db, err := sql.Open("pgx", "host=localhost port=5432 user=testuser dbname=testdb password=pass sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}
