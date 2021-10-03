package day6

import (
	"container/list"
	"runtime"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Queue struct {
	Val []*TreeNode
}

func (this *Queue) get() *TreeNode {
	if len(this.Val) == 0 {
		return nil
	}
	ret := this.Val[0]
	this.Val = this.Val[1:]
	return ret
}

func (this *Queue) push(x *TreeNode) {
	this.Val = append(this.Val, x)
}

func (this *Queue) isNil() bool {
	if len(this.Val) == 0 {
		return true
	}
	return false
}

var que Queue

func levelOrderV1(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	que.push(root)
	for !que.isNil() {
		node := que.get()
		// fmt.Println(node)
		res = append(res, node.Val)
		if node.Left != nil {
			que.push(node.Left)
		}
		if node.Right != nil {
			que.push(node.Right)
		}
	}
	return res
}

func levelOrderV2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	queue := list.New()
	queue.PushBack(root) // push
	ans := []int{}

	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode) // pop
		ans = append(ans, node.Val)
		if node.Left != nil {
			queue.PushBack(node.Left) // push
		}
		if node.Right != nil {
			queue.PushBack(node.Right) //push
		}
		runtime.GC()
	}

	return ans
}

func LevelOrder(root *TreeNode) []int {
	return levelOrderV2(root)
}

func LevelOrderV2(root *TreeNode) []int {
	return levelOrderV2(root)
}
