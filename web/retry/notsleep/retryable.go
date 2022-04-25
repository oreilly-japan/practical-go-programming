package notsleep

import (
	"net/http"
	"time"
)

// retryable-retry-start
type retryableRoundTripper struct {
	tr       http.RoundTripper
	attempts int
	waitTime time.Duration
}

func (t *retryableRoundTripper) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	for count := 0; count < t.attempts; count++ {
		resp, err = t.tr.RoundTrip(req)

		// リトライできない場合はリトライしない
		if err != nil || resp.StatusCode != http.StatusTooManyRequests {
			return resp, err
		}

		select {
		case <-req.Context().Done():
			return nil, req.Context().Err()
		case <-time.After(t.waitTime): // リトライのために待機
		}
	}
	return resp, err
}

// retryable-retry-end
