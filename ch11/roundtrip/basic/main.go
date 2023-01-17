package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

// basic-rt-start
type customRoundTripper struct {
	base http.RoundTripper
}

func (c customRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// << リクエスト前の実施したい処理を追加 >>
	resp, err := c.base.RoundTrip(req)
	// << リクエスト後に実施した処理を追加 >>
	return resp, err
}

// basic-rt-end

func main() {
	ctx := context.Background()
	// roundtrip-client-start
	client := &http.Client{
		Transport: &customRoundTripper{
			base: http.DefaultTransport,
		},
	}
	req, err := http.NewRequestWithContext(ctx, "GET", "http://example.com", nil)
	if err != nil {
		// エラーハンドリング
	}
	resp, err := client.Do(req)
	if err != nil {
		// ...
	}
	// roundtrip-client-end

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// ...
	}
	fmt.Println(string(body))
}
