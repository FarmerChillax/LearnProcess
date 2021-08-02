package main

import "fmt"

// demo
// func levelOrder(root *TreeNode) [][]int {
// 	ret := [][]int{}
// 	if root == nil {
// 		return ret
// 	}
// 	q := []*TreeNode{root}
// 	for i := 0; len(q) > 0; i++ {
// 		ret = append(ret, []int{})
// 		p := []*TreeNode{}
// 		for j := 0; j < len(q); j++ {
// 			node := q[j]
// 			ret[i] = append(ret[i], node.Val.(int))
// 			if node.Left != nil {
// 				p = append(p, node.Left)
// 			}
// 			if node.Right != nil {
// 				p = append(p, node.Right)
// 			}
// 		}
// 		q = p
// 	}
// 	return ret
// }

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val.(int))
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}

func main() {
	inputList := []nodeType{3, 9, 20, nil, nil, 15, 7}
	res := NewTree(&inputList)
	levelOrder(res)
	// fmt.Println(res)
	// res.postOrder()
	// queue := Queue{}
	// fmt.Println(queue)
	// queue.Add(123)
	// fmt.Println(queue)
	// fmt.Println(queue.isEmpty())
	// fmt.Println(queue.Remove())
	// fmt.Println(queue)
	// fmt.Println(queue.isEmpty())
}

// var queue Queue

type Queue []*TreeNode

func (q *Queue) Add(data *TreeNode) bool {
	*q = append(*q, data)
	return true
}

func (q *Queue) Remove() *TreeNode {
	node := (*q)[0]
	*q = (*q)[1:]
	return node
}

func (q *Queue) isEmpty() bool {
	if len(*q) == 0 {
		return true
	}
	return false
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

func (t *TreeNode) postOrder() {
	if t == nil || t.Val == nil {
		return
	}
	t.Left.postOrder()
	t.Right.postOrder()
	fmt.Println(t.Val)
}
