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
	if res == nil {
		fmt.Println(res)
	}
}

func Day6T2(t *testing.T) {
	binaryTree := []int{3, 9, 0, 0, 20, 15, 0, 0, 7, 0, 0}
	root := day6.NewTree(&binaryTree)
	res := day6.LevelOrderT2(root)
	fmt.Println(res)
}

func Day6T3(t *testing.T) {
	binaryTree := []int{3, 9, 0, 0, 20, 15, 0, 0, 7, 0, 0}
	root := day6.NewTree(&binaryTree)
	res := day6.LevelOrderT3(root)
	fmt.Println(res)
}

func TestDay6Sub(t *testing.T) {
	t.Run("day6=t1", Day6T1)
	t.Run("day6=t2", Day6T2)
	t.Run("day6=t3", Day6T3)
}
