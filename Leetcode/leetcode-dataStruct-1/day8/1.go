package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	node := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return node
}

func main() {
	inputList := []int{1, 2, 3, 4, 5, 6, 7, 8}

	pointer := NewListNode(inputList)
	// pointer.Print()
	res := reverseList(pointer)
	fmt.Println(res)
	// res.Print()
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
