package sort_test

import (
	"algorithrn/sorts"
	"testing"
)

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sorts.QuickSort([]int{12, 87, 1, 66, 38, 213, 126, 1233, 6542, 456, 973, 102, 3})
	}
}

func BenchmarkSelectSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sorts.SelectionSort([]int{12, 87, 1, 66, 38, 213, 126, 1233, 6542, 456, 973, 102, 3})
	}
}

func BenchmarkOptBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sorts.OptimizationBubbleSort([]int{12, 87, 1, 66, 38, 213, 126, 1233, 6542, 456, 973, 102, 3})
	}
}
func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sorts.BubbleSort([]int{12, 87, 1, 66, 38, 213, 126, 1233, 6542, 456, 973, 102, 3})
	}
}
