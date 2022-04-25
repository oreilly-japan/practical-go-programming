package main

import (
	"fmt"
	"log"
	"net/http"

	"go.uber.org/zap"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world!")
	}
	http.HandleFunc("/test", handler)

	// アプリケーションで使用するサードパーティのロガー
	l, err := zap.NewDevelopment()
	if err != nil {
		l = zap.NewNop()
	}
	logger := l.Sugar()

	server := &http.Server{
		Addr: ":18888",
		// io.Writeを満たす構造体でラップ
		ErrorLog: log.New(&logForwarder{l: logger}, "", 0),
	}

	logger.Fatal("server: %v", server.ListenAndServe())
}

// logForwarderは io.Writer インタフェースを満たすラッパー構造体
type logForwarder struct {
	l *zap.SugaredLogger
}

// ロガーの出力をWriteに移譲してWriteを実装することで io.Write を満たします
func (fw *logForwarder) Write(p []byte) (int, error) {
	fw.l.Errorw(string(p))
	return len(p), nil
}
