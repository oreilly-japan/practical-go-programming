package goretryablehttp

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
)

func hoge() {
	ctx := context.Background()

	// go-retryablehttp-start
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 2
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://example.com/", nil)
	resp, err := retryClient.StandardClient().Do(req)
	if err != nil {
		// ...
	}

	// go-retryablehttp-end
	defer resp.Body.Close()
	_ = resp
}
