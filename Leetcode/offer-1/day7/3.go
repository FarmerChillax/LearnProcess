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

func isSymmetric(root *tree.TreeNode) bool {
	return check(root, root)
}

func check(r1, r2 *tree.TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	}
	if r1 == nil || r2 == nil {
		return false
	}
	return r1.Val == r2.Val && check(r1.Left, r2.Right) && check(r1.Right, r2.Left)
}

func IsSymmetric(root *tree.TreeNode) bool {
	return isSymmetric(root)
}
