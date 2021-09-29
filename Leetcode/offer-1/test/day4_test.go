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
