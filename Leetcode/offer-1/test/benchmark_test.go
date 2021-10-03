package test

import (
	"offer/day2"
	"offer/day3"
	"offer/day4"
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

func BenchmarkDay3Sub1V1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = day3.ReplaceSpaceV1("We are happy.")
	}
}

func BenchmarkDay3Sub1V2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = day3.ReplaceSpaceV2("We are happy.")
	}
}

func BenchmarkDay3Sub1V3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = day3.ReplaceSpaceV3("We are happy.")
	}
}

func BenchmarkDay3Sub2V1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = day3.ReverseLeftWords("lrloseumgh", 6)
	}
}

func BenchmarkDay3Sub2V2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = day3.ReverseLeftWordsV2("lrloseumgh", 6)
	}
}

func BenchmarkDay4Sub1V1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = day4.FindRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3})
	}
}

func BenchmarkDay4Sub1V2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = day4.FindRepeatNumberV2([]int{2, 3, 1, 0, 2, 5, 3})
	}
}

func BenchmarkDay4Sub2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = day4.Search([]int{5, 7, 7, 8, 8, 10}, 8)
	}
}

func BenchmarkDay4Sub2V2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = day4.SearchV2([]int{5, 7, 7, 8, 8, 10}, 8)
	}
}
