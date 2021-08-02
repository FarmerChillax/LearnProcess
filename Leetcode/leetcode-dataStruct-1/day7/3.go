package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	root := &ListNode{Next: head}
	node := root
	for {
		if node == nil {
			break
		}
		if node.Next != nil && node.Next.Val == val {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}
	return root.Next
}
func main() {
	inputList := []int{1, 2, 6, 3, 4, 5, 6}
	val := 6
	point := NewListNode(inputList)
	res := removeElements(point, val)
	res.Print()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) Print() {
	node := n
	for {
		if node == nil {
			return
		}
		fmt.Println(node.Val)
		node = node.Next
	}
}

func NewListNode(inputList []int) *ListNode {
	head := &ListNode{}
	node := head
	for _, item := range inputList {
		tmp := ListNode{Val: item, Next: nil}
		node.Next = &tmp
		node = node.Next
	}
	return head.Next
}
