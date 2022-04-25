package cookbooktest

import "testing"

// in-start

func TestMain(m *testing.M) {
	// 1. テスト全体の実行前
	setup()
	// 6. テスト全体の実行後
	defer teardown()

	m.Run()
}

func setup() {
	// setup code
}

func teardown() {
	// tear down code
}

func TestHoge(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// テストケース
	}

	// 2. テスト関数の実行前
	defer func() {
		// 5. テスト関数の実行後
	}()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// 3. テストケースの実行前
			defer func() {
				// 4. テストケースの実行後
			}()

			got := Hoge(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("Hoge() = %v, want %v", got, tt.want)
			}

		})
	}
}

// in-end
