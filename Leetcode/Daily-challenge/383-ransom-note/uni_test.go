package ransomnote

import (
	"log"
	"reflect"
	"testing"
)

func TestCanConstruct_1(t *testing.T) {
	ransomNote := "a"
	magazine := "b"
	expected := false
	result := canConstruct(ransomNote, magazine)
	if !reflect.DeepEqual(result, expected) {
		log.Fatalln("not match, result:", result)
	}
}

func TestCanConstruct_2(t *testing.T) {
	ransomNote := "aa"
	magazine := "ab"
	expected := false
	result := canConstruct(ransomNote, magazine)
	if !reflect.DeepEqual(result, expected) {
		log.Fatalln("not match, result:", result)
	}
}

func TestCanConstruct_3(t *testing.T) {
	ransomNote := "aa"
	magazine := "aab"
	expected := false
	result := canConstruct(ransomNote, magazine)
	if !reflect.DeepEqual(result, expected) {
		log.Fatalln("not match, result:", result)
	}
}
