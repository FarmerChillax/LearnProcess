package test

import (
	"offer/day8"
	"testing"
)

func Day8T1(t *testing.T) {
	n := 1
	expected := 1
	res := day8.Fib(n)

	if expected != res {
		t.Errorf("not match...%v\n", res)
	}
}

func Day8T2(t *testing.T) {
	n := 7
	expected := 21
	res := day8.NumWays(n)

	if expected != res {
		t.Errorf("not match...%v\n", res)
	}
}

func Day8T3(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	expected := 5
	res := day8.MaxProfit(prices)

	if expected != res {
		t.Errorf("not match...%v\n", res)
	}
}

func Day8T3_2(t *testing.T) {
	prices := []int{7, 6, 4, 3, 1}
	expected := 0
	res := day8.MaxProfit(prices)

	if expected != res {
		t.Errorf("not match...%v\n", res)
	}
}
func TestDay8Sub(t *testing.T) {
	t.Run("day8=t1", Day8T1)
	t.Run("day8=t2", Day8T2)
	t.Run("day8=t3", Day8T3)
	t.Run("day8=t3-2", Day8T3_2)
}
