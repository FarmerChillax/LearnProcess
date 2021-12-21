package day3

import "testing"

func BenchmarkTestReplace(b *testing.B) {
	s := "We are happy."
	for i := 0; i < b.N; i++ {
		_ = replaceSpaceV1(s)
	}
}

func BenchmarkTestReplaceV2(b *testing.B) {
	s := "We are happy."
	for i := 0; i < b.N; i++ {
		_ = replaceSpaceV2(s)
	}
}

func BenchmarkTestReplaceV3(b *testing.B) {
	s := "We are happy."
	for i := 0; i < b.N; i++ {
		_ = replaceSpaceV3(s)
	}
}

// 第二题
func BenchmarkTestReverse(b *testing.B) {
	s := "We are happy."
	k := 2
	for i := 0; i < b.N; i++ {
		_ = reverseLeftWords(s, k)
	}
}

func BenchmarkTestReverseV2(b *testing.B) {
	s := "We are happy."
	k := 2
	for i := 0; i < b.N; i++ {
		_ = reverseLeftWordsV2(s, k)
	}
}
