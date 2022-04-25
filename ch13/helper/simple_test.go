package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// buffer-test
// テスト対象
type User struct {
	Name    string
	Address string
}

func DumpUser(u *User) {
	DumpUserTo(os.Stdout, u)
}

func DumpUserTo(w io.Writer, u *User) {
	if u.Address == "" {
		fmt.Fprintf(w, "%s(住所不定)", u.Name)
	} else {
		fmt.Fprintf(w, "%s@%s", u.Name, u.Address)
	}
}

// テストコード
func TestConsoleOut(t *testing.T) {
	var b bytes.Buffer
	DumpUserTo(&b, &User{Name: "クライド・バロウ"})
	if b.String() != "クライド・バロウ(住所不定)" {
		t.Errorf("error (expected: 'クライド・バロウ(住所不定)', actual='%s'", b.String())
	}
}

// buffer-test

// http-test
// テスト用ヘルパー
type testTransport struct {
	req **http.Request
	res *http.Response
	err error
}

func (t *testTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	*(t.req) = req
	return t.res, t.err
}

var _ http.RoundTripper = &testTransport{}

func newTransport(req **http.Request, res *http.Response, err error) http.RoundTripper {
	return &testTransport{
		req: req,
		res: res,
		err: err,
	}
}

// テストコード
func TestHTTPRequest(t *testing.T) {
	var req *http.Request
	res := httptest.NewRecorder()
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.WriteString(`{"ranking": ["Back to the Future", "Rambo"]}`)

	c := http.Client{
		Transport: newTransport(&req, res.Result(), nil),
	}
	r, err := c.Get("http://example.com/movies/1985")
	assert.Nil(t, err)
	assert.Equal(t, 200, r.StatusCode)
}
// http-test

// timer-test

// テスト用ヘルパー群
type contextTimeKey string

const timeKey contextTimeKey = "timeKey"

func CurrentTime(ctx context.Context) time.Time {
	v := ctx.Value(timeKey)
	if t, ok := v.(time.Time); ok {
		return t
	}
	return time.Now()
}

func SetFixTime(ctx context.Context, t time.Time) context.Context {
	return context.WithValue(ctx, timeKey, t)
}

// テスト対象
func NextMonth(ctx context.Context) time.Month {
	now := CurrentTime(ctx)
	return now.AddDate(0, 1, 0).Month()
}

// テストコード
func TestNextMonth(t *testing.T) {
	ctx := SetFixTime(context.Background(), time.Date(1980, time.December, 1, 0, 0, 0, 0, time.Local))
	assert.Equal(t, time.January, NextMonth(ctx))
}

// timer-test
