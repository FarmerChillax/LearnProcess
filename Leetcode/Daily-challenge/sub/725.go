package sub

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func makeList(valus []int) *ListNode {
	head := &ListNode{}
	node := head
	for _, v := range valus {
		node.Next = &ListNode{Val: v, Next: nil}
		node = node.Next
	}
	return head.Next
}

func (l *ListNode) print() {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	listLen := 0
	for node := head; node != nil; node = node.Next {
		listLen++
	}

	aveLength, remainder := listLen/k, listLen%k
	parts := make([]*ListNode, k)

	for i, node := 0, head; i < k && node != nil; i++ {
		parts[i] = node

		partSize := aveLength
		if i < remainder {
			partSize++
		}

		for j := 1; j < partSize; j++ {
			node = node.Next
		}
		nowPoint := node
		node = node.Next
		nowPoint.Next = nil
	}
	return parts
}

// func main() {
// 	headList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	head := makeList(headList)
// 	// head.print()
// 	res := splitListToParts(head, 3)
// 	fmt.Println(res[0])
// }

// func splitListToParts(head *ListNode, k int) []*ListNode {
// 	n := 0
// 	for node := head; node != nil; node = node.Next {
// 		n++
// 	}
// 	quotient, remainder := n/k, n%k

// 	parts := make([]*ListNode, k)
// 	for i, curr := 0, head; i < k && curr != nil; i++ {
// 		parts[i] = curr
// 		partSize := quotient
// 		if i < remainder {
// 			partSize++
// 		}
// 		for j := 1; j < partSize; j++ {
// 			curr = curr.Next
// 		}
// 		curr, curr.Next = curr.Next, nil
// 	}
// 	return parts
// }
