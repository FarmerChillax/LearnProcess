package main

import (
	"fmt"
)

type nodeType interface{}

type TreeNode struct {
	value       nodeType
	left, right *TreeNode
}

// BinaryTree proOrder function
func (t *TreeNode) preOrder() {
	if t == nil || t.value == nil {
		return
	}
	fmt.Println(t.value)
	t.left.preOrder()
	t.right.preOrder()
}

// binaryTree inOrder function
func (t *TreeNode) inOrder() {
	if t == nil || t.value == nil {
		return
	}
	t.left.inOrder()
	fmt.Println(t.value)
	t.right.inOrder()
}

// BinaryTree postOrder function
func (t *TreeNode) postOrder() {
	if t == nil || t.value == nil {
		return
	}
	t.left.postOrder()
	t.right.postOrder()
	fmt.Println(t.value)

}

func newTree(inputList *[]nodeType) *TreeNode {
	if len(*inputList) == 0 || inputList == nil {
		return nil
	}

	data := (*inputList)[0]
	*inputList = (*inputList)[1:]

	if data == nil {
		return nil
	}

	node := &TreeNode{data, nil, nil}
	node.left = newTree(inputList)
	node.right = newTree(inputList)
	return node
}

func main() {
	inputList := []nodeType{3, 2, 9, nil, nil, 10, nil, nil, 8, nil, 4}
	root := newTree(&inputList)
	fmt.Println("res", root.left.right.value)
	fmt.Println("前序遍历")
	root.preOrder()
	fmt.Println("中序遍历")
	root.inOrder()
	fmt.Println("后续遍历")
	root.postOrder()
}
