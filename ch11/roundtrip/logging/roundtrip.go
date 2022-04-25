package logging

import (
	"log"
	"net/http"
	"time"
)

func NewAppRetryableClient() *http.Client {
	return &http.Client{
		Transport: &loggingRoundTripper{transport: http.DefaultTransport},
	}
}

// start-roundtrip
type loggingRoundTripper struct {
	transport http.RoundTripper
	logger    func(string, ...interface{})
}

func (t *loggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.logger == nil {
		t.logger = log.Printf
	}

	start := time.Now()
	resp, err := t.transport.RoundTrip(req)

	if resp != nil {
		t.logger("%s %s %d %s, duration: %d",
			req.Method, req.URL.String(), resp.StatusCode, http.StatusText(resp.StatusCode), time.Since(start))
	}

	return resp, err
}

// end-roundtrip
