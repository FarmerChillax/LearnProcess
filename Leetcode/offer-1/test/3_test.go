package test

import (
	"offer/day3"
	"reflect"
	"testing"
)

func TestDay3Sub1(t *testing.T) {
	s := "We are happy."
	expected := "We%20are%20happy."
	res := day3.ReplaceSpace(s)
	if res != expected {
		t.Errorf("not match, res: %v\n", res)
	}
}

func TestDay3Sub2(t *testing.T) {
	s := "abcdefg"
	k := 2
	expected := "cdefgab"
	res := day3.ReverseLeftWords(s, k)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("not match, res: %v\n", res)
	}
}

func TestDay3Sub2_1(t *testing.T) {
	s := "lrloseumgh"
	k := 6
	expected := "umghlrlose"
	res := day3.ReverseLeftWords(s, k)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("not match, res: %v\n", res)
	}
}

func TestDay3Sub2V2(t *testing.T) {
	s := "abcdefg"
	k := 2
	expected := "cdefgab"
	res := day3.ReverseLeftWordsV2(s, k)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("not match, res: %v\n", res)
	}
}

func TestDay3Sub2V2_1(t *testing.T) {
	s := "lrloseumgh"
	k := 6
	expected := "umghlrlose"
	res := day3.ReverseLeftWordsV2(s, k)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("not match, res: %v\n", res)
	}
}
