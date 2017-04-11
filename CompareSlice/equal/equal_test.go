package equal

import "testing"

func BenchmarkEqual(b *testing.B) {
	sa := []string{"q", "w", "e", "r", "t"}
	sb := []string{"q", "w", "a", "s", "z", "x"}
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		StringSliceEqual(sa, sb)
	}
}

func BenchmarkDeepEqual(b *testing.B) {
	sa := []string{"q", "w", "e", "r", "t"}
	sb := []string{"q", "w", "a", "s", "z", "x"}
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		StringSliceReflectEqual(sa, sb)
	}
}
