package longtest

import (
	"testing"
)

// skip-test
func TestTimeConsuming(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	// 時間がかかるテスト
}

// skip-test
