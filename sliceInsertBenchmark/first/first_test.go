package first

import "testing"

func BenchmarkFirst(b *testing.B) {
	// run the First function b.N times
	for n := 0; n < b.N; n++ {
		First()
	}
}
