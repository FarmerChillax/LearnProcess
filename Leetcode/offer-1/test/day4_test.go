package test

import (
	"offer/day4"
	"testing"
)

func TestSub1V1(t *testing.T) {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	res := day4.FindRepeatNumber(nums)
	if res != 3 && res != 2 {
		t.Errorf("result err, return res is: %v\n", res)
	}
}

func TestSub1V2(t *testing.T) {
	nums := []int{3, 1, 2, 3}
	res := day4.FindRepeatNumberV2(nums)
	if res != 3 {
		t.Errorf("result err, return res is: %v\n", res)
	}
}

func TestSub2V1(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	expected := 2
	res := day4.Search(nums, target)
	if res != expected {
		t.Errorf("result err, return res is: %v\n", res)
	}
}

func TestSub2_2(t *testing.T) {
	nums := []int{1}
	target := 1
	expected := 1
	res := day4.Search(nums, target)
	if res != expected {
		t.Errorf("result err, return res is: %v\n", res)
	}
}

func TestSub2V2(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	expected := 2
	res := day4.SearchV2(nums, target)
	if res != expected {
		t.Errorf("result err, return res is: %v\n", res)
	}
}
func TestSub3(t *testing.T) {
	nums := []int{0}
	expected := 1
	res := day4.MissingNumber(nums)
	if res != expected {
		t.Errorf("result err, return res is: %v\n", res)
	}
}
