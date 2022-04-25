package main

import (
	"testing"
)

func TestLog(t *testing.T) {
	t.Log("失敗時やgo test -vのときだけ表示されます")
	t.Fatal("メッセージとともにテストを失敗させます")
	// --- FAIL: TestLog (0.00s)
	//     log_test.go:8: 失敗時やgo test -vのときだけ表示されます
	//     log_test.go:9: メッセージとともにテストを失敗させます
}
