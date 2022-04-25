package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"
)

func threadlocal() {
	// sample:context-basic
	ctx := context.Background() // ルートのcontext
	ctx = context.WithValue(ctx, "key", "value")
	log.Println(ctx.Value("key"))
	// sample:context-basic
}

func uniqueKey() {
	var contextKey = struct{}{} // 他と重複しないcontextキー
	ctx := context.Background()
	ctx = context.WithValue(ctx, contextKey, "value")
	log.Println(ctx.Value(contextKey))
}

// sample:store-for-context
var tokenContextKey = struct{}{}

func RegisterToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, tokenContextKey, token)
}

func RetrieveToken(ctx context.Context) (string, error) {
	token, ok := ctx.Value(tokenContextKey).(string)
	if !ok {
		return "", errors.New("Token is not registered")
	}
	return token, nil
}

// sample:store-for-context

// sample:store-for-req
func RegisterToken2(req *http.Request, token string) *http.Request {
	ctx := context.WithValue(req.Context(), tokenContextKey, token)
	return req.WithContext(ctx)
}

func RetrieveToken2(req *http.Request) (string, error) {
	token, ok := req.Context().Value(tokenContextKey).(string)
	if !ok {
		return "", errors.New("Token is not registered")
	}
	return token, nil
}

// sample:store-for-req

// sample:store-log
var logContextKey = struct{}{}

type LogRecord struct {
	start         time.Time
	Duration      time.Duration
	DBAccessCount int
}

func StartLogging(req *http.Request) *http.Request {
	l := &LogRecord{
		start: time.Now(),
	}
	ctx := context.WithValue(req.Context(), logContextKey, l)
	return req.WithContext(ctx)
}

func CountAccess(ctx context.Context) error {
	l, ok := ctx.Value(logContextKey).(*LogRecord)
	if !ok {
		return errors.New("Logger is not registered")
	}
	l.DBAccessCount += 1
	return nil
}

func FinishLogging(req *http.Request) (*LogRecord, error) {
	l, ok := req.Context().Value(logContextKey).(*LogRecord)
	if !ok {
		return nil, errors.New("Logger is not registered")
	}
	l.Duration = time.Now().Sub(l.start)
	return l, nil
}

// sample:store-log
