package second

import "testing"

func BenchmarkSecond(b *testing.B) {
	// run the Second function b.N times
	for n := 0; n < b.N; n++ {
		Second()
	}
}
