package detectleak

import (
	"testing"

	"go.uber.org/goleak"
)

// setting
func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

// setting

// leaktest
func TestLeak(t *testing.T) {
	wait := make(chan struct{})
	go func() {
		// ここでブロックするのでリーク
		<-wait
	}()
}

// leaktest
