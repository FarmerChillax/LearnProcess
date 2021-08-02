package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//  my first function
// func deleteDuplicates(head *ListNode) *ListNode {
// 	numMap := map[int]int{}
// 	node := head
// 	for {
// 		if node == nil {
// 			break
// 		}
// 		numMap[node.Val]++
// 		node = node.Next
// 	}
// 	node = head
// 	for {
// 		if node == nil {
// 			break
// 		}
// 		if node.Next == nil {
// 			break
// 		}
// 		next := node.Next
// 		if numMap[next.Val] > 1 {
// 			node.Next = next.Next
// 			numMap[next.Val]--
// 		} else {
// 			node = node.Next
// 		}
// 	}

// 	return head
// }

// one for func
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	node := head
	for node.Next != nil {
		if node.Val == node.Next.Val {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}
	return head
}

func main() {
	inputList := []int{1, 1, 1, 2, 3, 4, 5}

	pointer := NewListNode(inputList)
	res := deleteDuplicates(pointer)
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
