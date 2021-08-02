package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root == p || root == q {
		return root
	}
	var leftNode, rightNode *TreeNode

	if p.Val.(int) > q.Val.(int) {
		// 左 大于 右 则交换。保证左边小于右边（p < q)
		p, q = q, p
	}
	leftNode = lowestCommonAncestor(root.Left, p, q)
	rightNode = lowestCommonAncestor(root.Right, p, q)
	if leftNode != nil && rightNode != nil {
		return root
	}
	if leftNode != nil {
		return leftNode
	}
	if rightNode != nil {
		return rightNode
	}
	return nil
}

func main() {
	inputList := []nodeType{6, 2, 0, -1, -1, 4, 3, -1, -1, 5, -1, -1, 8, 7, -1, -1, 9}
	root := NewTree(&inputList)
	fmt.Println(root)
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

	if data == -1 {
		return nil
	}

	node := &TreeNode{data, nil, nil}
	node.Left = NewTree(inputList)
	node.Right = NewTree(inputList)
	return node
}
