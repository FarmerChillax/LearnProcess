package sub_test

import (
	"daily-challenge/sub"
	"testing"
)

func BenchmarkSub69(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_ = sub.PeakIndexInMountainArray([]int{3, 4, 5, 1})
	}
}
func BenchmarkSub69V2(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_ = sub.PeakIndexInMountainArrayV2([]int{3, 4, 5, 1})
	}
}
