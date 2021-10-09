package test

import (
	"fmt"
	"offer/day7"
	"offer/tree"
	"testing"
)

func Day7T1(t *testing.T) {
	binaryTree := []int{3, 4, 1, 0, 0, 2, 0, 0, 5, 0, 0}
	secTree := []int{4, 1, 0, 0, 0}
	root := tree.NewTree(&binaryTree)
	secRoot := tree.NewTree(&secTree)
	res := day7.IsSubStructure(root, secRoot)
	fmt.Println(res)
}

func Day7T2(t *testing.T) {
	binaryTree := []int{3, 9, 0, 0, 20, 15, 0, 0, 7, 0, 0}
	root := tree.NewTree(&binaryTree)
	res := day7.MirrorTree(root)
	fmt.Println(res.Left)
}

func Day7T3(t *testing.T) {
	binaryTree := []int{1, 2, 3, 0, 0, 4, 0, 0, 2, 4, 0, 0, 3, 0, 0}
	root := tree.NewTree(&binaryTree)
	res := day7.IsSymmetric(root)
	fmt.Println(res)
}

func TestDay7Sub(t *testing.T) {
	t.Run("day7=t2", Day7T2)
	t.Run("day7=t3", Day7T3)
	t.Run("day7=t1", Day7T1)
}
