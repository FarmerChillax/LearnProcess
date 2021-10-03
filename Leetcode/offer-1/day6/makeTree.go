package day6

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func NewTree(inputList *[]int) *TreeNode {
	if len(*inputList) == 0 || inputList == nil {
		return nil
	}

	data := (*inputList)[0]
	*inputList = (*inputList)[1:]
	// fmt.Println(data)
	if data == 0 {
		return nil
	}

	node := &TreeNode{data, nil, nil}
	node.Left = NewTree(inputList)
	node.Right = NewTree(inputList)
	return node
}

func (t *TreeNode) PostOrder() {
	if t == nil || t.Val == 0 {
		return
	}
	t.Left.PostOrder()
	t.Right.PostOrder()
	fmt.Println(t.Val)

}

func (t *TreeNode) PreOrder() {
	if t == nil {
		return
	}
	fmt.Println(t.Val)
	t.Left.PreOrder()

	t.Right.PreOrder()
}
