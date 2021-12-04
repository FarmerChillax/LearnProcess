package day1

import (
	"fmt"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	raw := []int{2, 2, 1}
	res := singleNumber(raw)
	fmt.Println(res)
}

func TestMajorityElement(t *testing.T) {
	raw := []int{3, 2, 3}
	res := majorityElement(raw)
	fmt.Println(res)
}

func TestThreeSum(t *testing.T) {
	raw := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(raw)
	fmt.Println(res)
}
