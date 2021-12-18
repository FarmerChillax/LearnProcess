package wordpattern

import (
	"log"
	"testing"
)

func TestWordPattern_1(t *testing.T) {
	pattern := "abba"
	str := "dog cat cat dog"
	excepted := true
	result := wordPattern(pattern, str)
	if excepted != result {
		log.Fatalln("not match, result:", result)
	}
}

func TestWordPattern_2(t *testing.T) {
	pattern := "abba"
	str := "dog cat cat fish"
	excepted := false
	result := wordPattern(pattern, str)
	if excepted != result {
		log.Fatalln("not match, result:", result)
	}
}

func TestWordPattern_3(t *testing.T) {
	pattern := "aaaa"
	str := "dog cat cat dog"
	excepted := false
	result := wordPattern(pattern, str)
	if excepted != result {
		log.Fatalln("not match, result:", result)
	}
}

func TestWordPattern_4(t *testing.T) {
	pattern := "abba"
	str := "dog dog dog dog"
	excepted := false
	result := wordPattern(pattern, str)
	if excepted != result {
		log.Fatalln("not match, result:", result)
	}
}
