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
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDeep := maxDepth(root.Left) + 1
	rightDeep := maxDepth(root.Right) + 1
	if leftDeep >= rightDeep {
		return leftDeep
	} else {
		return rightDeep
	}
}

func main() {
	inputList := []nodeType{3, 2, 9, nil, nil, 10, nil, nil, 8, nil, 4}
	root := NewTree(&inputList)
	res := maxDepth(root)
	fmt.Println(res)
}

type nodeType interface{}

type TreeNode struct {
	Val         nodeType
	Left, Right *TreeNode
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
