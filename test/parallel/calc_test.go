package cookbooktest

import (
	"testing"
)

// in-start
func TestContains(t *testing.T) {
	type args struct {
		src []string
		dst string
	}
	tests := []struct {
		name   string
		args   args
		want   int
		hasErr bool
	}{
		{
			name: "対象が含まれる",
			args: args{
				src: []string{"a", "b", "c"},
				dst: "b",
			},
			want:   1,
			hasErr: false,
		},
		{
			name: "対象が含まれない",
			args: args{
				src: []string{"a", "b", "c"},
				dst: "d",
			},
			hasErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Contains(tt.args.src, tt.args.dst)
			if (err != nil) != tt.hasErr {
				t.Errorf("unexpected error: %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

// in-end

func TestCalc(t *testing.T) {
	type args struct {
		a        int
		b        int
		operator string
	}
	tests := []struct {
		name   string
		args   args
		want   int
		hasErr bool
	}{
		{
			name: "足し算",
			args: args{
				a:        10,
				b:        2,
				operator: "+",
			},
			want:   12,
			hasErr: false,
		},
		{
			name: "不正な演算子を指定",
			args: args{
				a:        10,
				b:        2,
				operator: "?",
			},
			hasErr: true,
		},
	}

	// parallel-start
	for _, tt := range tests {
		tt := tt // 追加
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() // 追加
			got, err := Calc(tt.args.a, tt.args.b, tt.args.operator)
			// parallel-end

			if (err != nil) != tt.hasErr {
				t.Errorf("Calc() error = %v, hasErr %v", err, tt.hasErr)
				return
			}
			if got != tt.want {
				t.Errorf("Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}
