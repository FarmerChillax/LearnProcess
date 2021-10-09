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

func isSubStructure(A, B *tree.TreeNode) bool {
	res := false
	if A == nil || B == nil {
		return false
	}
	if A.Val == B.Val {
		res = subCheck(A, B)
	}
	return res || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func subCheck(ARoot, BRoot *tree.TreeNode) bool {

	if BRoot == nil {
		return true
	}

	if ARoot == nil {
		return false
	}

	if ARoot.Val == BRoot.Val {
		return subCheck(ARoot.Left, BRoot.Left) && subCheck(ARoot.Right, BRoot.Right)
	} else {
		return false
	}

}

func IsSubStructure(A, B *tree.TreeNode) bool {
	return isSubStructure(A, B)
}
