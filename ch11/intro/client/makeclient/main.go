package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	// client-start
	client := &http.Client{
		Timeout:   10 * time.Second, // タイムアウトを10秒に設定
		Transport: http.DefaultTransport,
	}
	// client-end

	// client-do-start
	// GETリクエストの生成
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://example.com", nil)
	if err != nil {
		// ...
	}
	// リクエストヘッダーにトークンを付与
	req.Header.Add("Authorization", "Bearer XXX...XXX")

	// HTTPリクエストの発行
	resp, err := client.Do(req)
	// client-do-end

	if err != nil {
		// ...
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// ...
	}
	fmt.Println(resp.Status)
	fmt.Println(string(body))
}
