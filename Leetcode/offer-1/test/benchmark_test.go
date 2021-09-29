package test

import (
	"offer/day2"
	"testing"
)

func BenchmarkDay2Sub1(b *testing.B) {
	lists := []int{1, 3, 2}
	head := makeList(lists)
	for i := 0; i < b.N; i++ {
		_ = day2.ReversePrint(head)
	}
}

func BenchmarkDay2Sub1V2(b *testing.B) {
	lists := []int{1, 3, 2}
	head := makeList(lists)
	for i := 0; i < b.N; i++ {
		_ = day2.ReversePrintV2(head)
	}
}

func BenchmarkDay2Sub1V3(b *testing.B) {
	lists := []int{1, 3, 2}
	head := makeList(lists)
	for i := 0; i < b.N; i++ {
		_ = day2.ReversePrintV3(head)
	}
}
