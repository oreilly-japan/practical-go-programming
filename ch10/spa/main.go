package main

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"time"

	"github.com/go-chi/chi/v5"
)

// assets-define
//go:embed vite-project/dist/*
var assets embed.FS
// assets-define

// try-read
func tryRead(requestedPath string, w http.ResponseWriter) error {
	log.Println(requestedPath)
	f, err := assets.Open(path.Join("vite-project/dist", requestedPath))
	if err != nil {
		return err
	}
	defer f.Close()

	// ディレクトリならエラー
	stat, _ := f.Stat()
	if stat.IsDir() {
		return errors.New("path is dir")
	}

	// MIMEタイプを設定
	ext := filepath.Ext(requestedPath)
	var contentType string
	if m := mime.TypeByExtension(ext); m != "" {
		contentType = m
	} else {
		contentType = "application/octet-stream"
	}
	w.Header().Set("Content-Type", contentType)
	io.Copy(w, f)
	return nil
}
// try-read

// other-implement
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	// まずはリクエストされた通りにファイルを探索
	err := tryRead(r.URL.Path, w)
	if err == nil {
		return
	}
	// 見つからなければindex.htmlを返す
	err = tryRead("index.html", w)
	if err != nil {
		panic(err)
	}
}

func newHandler() http.Handler {
	router := chi.NewRouter()

	router.Route("/api", func(r chi.Router) {
		// 何かAPIを足したい場合はここに足す
		r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		})
	})
	// シングルページアプリケーションを配布するハンドラー
	router.NotFound(notFoundHandler)

	return router
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	server := &http.Server{
		Addr:    ":8000" ,
		Handler: newHandler(),
	}

	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		server.Shutdown(ctx)
	}()
	fmt.Printf("start receiving at :8000")
	fmt.Fprintln(os.Stderr, server.ListenAndServe())
}
// other-implement
