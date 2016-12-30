package third

import "testing"

func BenchmarkThird(b *testing.B) {
	// run the Third function b.N times
	for n := 0; n < b.N; n++ {
		Third()
	}
}
