package cookbooktest

import "testing"

func BenchmarkAppendSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppendSlice(10000, 111)
	}
}

func BenchmarkFirstAllocSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FirstAllocSlice(10000, 111)
	}
}
