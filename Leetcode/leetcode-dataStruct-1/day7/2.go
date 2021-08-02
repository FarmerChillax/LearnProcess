package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	node := res
	for {
		if l1 == nil {
			node.Next = l2
			return res.Next
		}
		if l2 == nil {
			node.Next = l1
			return res.Next
		}
		tmp := &ListNode{}
		if l1.Val <= l2.Val {
			tmp.Val = l1.Val
			l1 = l1.Next
		} else {
			tmp.Val = l2.Val
			l2 = l2.Next
		}
		node.Next = tmp
		node = node.Next
	}
}

func main() {
	inputList1 := []int{1, 2, 4}
	inputList2 := []int{1, 3, 4}
	point1 := NewListNode(inputList1)
	point2 := NewListNode(inputList2)
	// fmt.Println(point1)
	point1.Print()
	fmt.Println(point2)
	res := mergeTwoLists(point1, point2)
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
