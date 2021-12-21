package day3

import (
	"log"
	"testing"
)

func TestReplace(t *testing.T) {
	s := "We are happy."
	expected := "We%20are%20happy."
	result := replaceSpaceV1(s)
	if result != expected {
		log.Fatalf("not match, result: %v; except: %v\n", result, expected)
	}
}

func TestReverse(t *testing.T) {
	s := "abcdefg"
	k := 2
	expected := "cdefgab"
	result := reverseLeftWordsV2(s, k)
	if result != expected {
		log.Fatalf("not match, result: %v; except: %v\n", result, expected)
	}
}
