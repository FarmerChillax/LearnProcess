package test

import (
	"offer/day5"
	"testing"
)

func T1Sub1(t *testing.T) {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	target := 5
	expected := true

	res := day5.FindNumberIn2DArray(matrix, target)
	if res != expected {
		t.Errorf("result error, res value: %v\n", res)
	}
}

func T1Sub2(t *testing.T) {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	target := 20
	expected := false

	res := day5.FindNumberIn2DArray(matrix, target)
	if res != expected {
		t.Errorf("result error, res value: %v\n", res)
	}
}

func T3Sub1(t *testing.T) {
	input := "z"
	expected := byte('z')
	res := day5.FirstUniqChar(input)
	if res != expected {
		t.Errorf("result error, res value: %c\n", res)
	}
}

func T2Sub1(t *testing.T) {
	input := []int{3, 4, 5, 1, 2}
	expected := 1
	res := day5.MinArray(input)
	if res != expected {
		t.Errorf("result error, res value: %c\n", res)
	}
}

func T2Sub2(t *testing.T) {
	input := []int{3, 5, 1}
	expected := 1
	res := day5.MinArray(input)
	if res != expected {
		t.Errorf("result error, res value: %v\n", res)
	}
}

func T2Sub3(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := 1
	res := day5.MinArray(input)
	if res != expected {
		t.Errorf("result error, res value: %c\n", res)
	}
}

func TestDay5Sub(t *testing.T) {
	// t.Run("sub1=1", T1Sub1)
	// t.Run("sub1=2", T1Sub2)
	t.Run("T3=1", T3Sub1)
	// t.Run("T2=1", T2Sub1)
	t.Run("T2=2", T2Sub2)
	// t.Run("T2=3", T2Sub3)
}
