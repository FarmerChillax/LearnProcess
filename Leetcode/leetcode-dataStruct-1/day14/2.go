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
func findTarget(root *TreeNode, k int) bool {
	hashMap := make(map[int]bool, 100*100)
	return preOrder(root, hashMap, k)
}

func preOrder(root *TreeNode, hashMap map[int]bool, k int) bool {
	if root == nil {
		return false
	}
	if _, ok := hashMap[k-root.Val.(int)]; ok {
		return true
	}
	hashMap[root.Val.(int)] = true
	return preOrder(root.Left, hashMap, k) || preOrder(root.Right, hashMap, k)
}

func main() {
	inputList := []nodeType{5, 3, 2, nil, nil, 4, nil, nil, 6, nil, 7, nil, nil}
	k := 9
	root := NewTree(&inputList)
	fmt.Println(root)
	fmt.Println(findTarget(root, k))
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
