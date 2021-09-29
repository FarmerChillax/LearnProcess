package day3

import "fmt"

func reverseLeftWords(s string, n int) string {
	return fmt.Sprintf("%s%s", s[n:], s[:n])
}

func ReverseLeftWords(s string, n int) string {
	return reverseLeftWords(s, n)
}

func reverseLeftWordsV2(s string, n int) string {
	return s[n:] + s[:n]
}

func ReverseLeftWordsV2(s string, n int) string {
	return reverseLeftWordsV2(s, n)
}
