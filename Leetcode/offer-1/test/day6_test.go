package test

import (
	"fmt"
	"offer/day6"
	"testing"
)

func Day6T1(t *testing.T) {
	inputList := []int{3, 2, 9, 0, 0, 10, 0, 0, 8, 0, 4}
	root := day6.NewTree(&inputList)
	res := day6.LevelOrder(root)
	fmt.Println(res)
}

func TestDay6Sub(t *testing.T) {
	t.Run("day6=t1", Day6T1)
}
