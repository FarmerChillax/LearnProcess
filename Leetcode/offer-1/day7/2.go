package day7

import "offer/tree"

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
 */

func mirrorTree(root *tree.TreeNode) *tree.TreeNode {
	root.Left, root.Right = root.Right, root.Left
	if root.Left != nil {
		mirrorTree(root.Left)
	}
	if root.Right != nil {
		mirrorTree(root.Right)
	}
	return root
}

func MirrorTree(root *tree.TreeNode) *tree.TreeNode {
	return mirrorTree(root)
}
