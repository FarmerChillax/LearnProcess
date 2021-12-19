package day1

import "testing"

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TestStackQueue(&testing.T{})
	}
}
