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
	rc := retryClient.StandardClient()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://example.com/", nil)
	// go-retryablehttp-end

	_, _, _ = rc, req, err
}
