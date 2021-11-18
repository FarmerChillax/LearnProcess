package sub

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findTilt(root *TreeNode) int {
	var dfs func(*TreeNode) int
	var res int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		sumLeft := dfs(node.Left)
		sumRight := dfs(node.Right)
		res += abs(sumLeft - sumRight)
		return sumLeft + sumRight + node.Val
	}
	dfs(root)
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func FindTilt(root *TreeNode) int {
	return findTilt(root)
}
