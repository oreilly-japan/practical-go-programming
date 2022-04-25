package retry

import (
	"errors"
	"net"
	"net/http"
	"time"
)

// retryable-retry-start
type retryableRoundTripper struct {
	base     http.RoundTripper
	attempts int
	waitTime time.Duration
}

func (rt *retryableRoundTripper) shouldRetry(resp *http.Response, err error) bool {
	// ネットワークエラーによるリトライ
	if err != nil {
		var netErr net.Error
		if errors.As(err, &netErr) && netErr.Temporary() {
			return true
		}
	}

	// レスポンスコードによるリトライ
	if resp != nil {
		if resp.StatusCode == 429 || (500 <= resp.StatusCode && resp.StatusCode <= 504) {
			return true
		}
	}

	// リトライすべきではないため、リトライしない
	return false
}

func (rt *retryableRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	var (
		resp *http.Response
		err  error
	)
	for count := 0; count < rt.attempts; count++ {
		resp, err = rt.base.RoundTrip(req)

		if !rt.shouldRetry(resp, err) {
			return resp, err
		}

		select {
		case <-req.Context().Done():
			return nil, req.Context().Err()
		case <-time.After(rt.waitTime): // リトライのために待機
		}
	}

	return resp, err
}

// retryable-retry-end
