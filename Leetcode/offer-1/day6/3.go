package day6

func levelOrderT3(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}
	queue := Queue{}
	queue.push(root)
	flag := 1
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
		if flag == -1 {
			tmp = reverse(tmp)
		}
		flag *= -1
		res = append(res, tmp)
	}
	return res
}

func reverse(revSlice []int) []int {
	for i, j := 0, len(revSlice)-1; i < j; i, j = i+1, j-1 {
		revSlice[i], revSlice[j] = revSlice[j], revSlice[i]
	}
	return revSlice
}

func LevelOrderT3(root *TreeNode) [][]int {
	return levelOrderT3(root)
}
