package day6

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrderT2(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}

	queue := Queue{}
	queue.push(root)
	for !queue.isNil() {
		var tmp []int
		var newQue []*TreeNode
		for _, node := range queue.Val {
			// fmt.Println(node)
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				newQue = append(newQue, node.Left)
			}
			if node.Right != nil {
				newQue = append(newQue, node.Right)
			}
		}
		queue.Val = newQue
		res = append(res, tmp)
	}
	return res
}

func LevelOrderT2(root *TreeNode) [][]int {
	return levelOrderT2(root)
}
