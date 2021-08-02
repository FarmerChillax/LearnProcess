package main

import (
	"fmt"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return validBST(root, math.MinInt64, math.MaxInt8)
}

func validBST(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val.(int) <= min || root.Val.(int) >= max {
		return false
	}
	return validBST(root.Left, min, root.Val.(int)) && validBST(root.Right, root.Val.(int), max)
}

func main() {
	inputList := []nodeType{2, 1, nil, nil, 3, nil, nil}
	inputList = []nodeType{5, 1, nil, nil, 4, 3, nil, nil, 6, nil, nil}
	inputList = []nodeType{2, 2, nil, nil, 2, nil, nil}
	root := NewTree(&inputList)
	fmt.Println(root)
	fmt.Println(isValidBST(root))
}

type nodeType interface{}

type TreeNode struct {
	Val   nodeType
	Left  *TreeNode
	Right *TreeNode
}

func NewTree(inputList *[]nodeType) *TreeNode {
	if len(*inputList) == 0 || inputList == nil {
		return nil
	}

	data := (*inputList)[0]
	*inputList = (*inputList)[1:]

	if data == nil {
		return nil
	}

	node := &TreeNode{data, nil, nil}
	node.Left = NewTree(inputList)
	node.Right = NewTree(inputList)
	return node
}
