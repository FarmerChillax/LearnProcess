package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return targetSum-root.Val.(int) == 0
	}

	leftStatus := hasPathSum(root.Left, targetSum-root.Val.(int))
	rightStatus := hasPathSum(root.Right, targetSum-root.Val.(int))

	return leftStatus || rightStatus
}

func main() {
	inputList := []nodeType{5, 4, 11, 7, nil, nil, 2, nil, nil, nil, 8, 13, nil, nil, 4, nil, 1, nil, nil}
	// inputList := []nodeType{1, 2, nil, nil, 3, nil, nil}
	targetSum := 22
	root := NewTree(&inputList)
	fmt.Println(hasPathSum(root, targetSum))
	// fmt.Println(root.Left)
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
