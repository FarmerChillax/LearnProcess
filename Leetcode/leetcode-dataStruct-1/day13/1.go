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
//  time: 99.54%, memory: 47.92%
// func searchBST(root *TreeNode, val int) *TreeNode {
// 	var node *TreeNode
// 	if root == nil {
// 		return node
// 	}
// 	if root.Val.(int) > val {
// 		node = searchBST(root.Left, val)
// 	}
// 	if root.Val.(int) < val {
// 		node = searchBST(root.Right, val)
// 	}
// 	if root.Val.(int) == val {
// 		node = root
// 	}
// 	return node
// }

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val.(int) == val {
		return root
	}
	if root.Val.(int) > val {
		return searchBST(root.Left, val)
	}
	if root.Val.(int) < val {
		return searchBST(root.Right, val)
	}
	return nil
}

func main() {
	inputList := &[]nodeType{4, 2, 1, nil, nil, 3, nil, nil, 7, nil, nil}
	target := 2
	root := NewTree(inputList)
	res := searchBST(root, target)
	fmt.Println(res)
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
