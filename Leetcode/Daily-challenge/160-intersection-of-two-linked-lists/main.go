// package intersectionoftwolinkedlists
package main

import "fmt"

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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	comparMap := map[*ListNode]bool{}
	for tmp := headA; tmp != nil; tmp = tmp.Next {
		comparMap[tmp] = true
	}

	for tmp := headB; tmp != nil; tmp = tmp.Next {
		if comparMap[tmp] {
			return tmp
		}
	}
	return nil
}

func generate(l []int) *ListNode {
	head := &ListNode{}
	point := head
	for _, item := range l {
		point.Next = &ListNode{Val: item}
		point = point.Next
	}
	return head.Next
}

func main() {
	headA := generate([]int{4, 1, 8, 4, 5})
	headB := generate([]int{5, 6, 1, 8, 4, 5})
	intersectVal := 8
	result := getIntersectionNode(headA, headB)
	fmt.Println(result, intersectVal)
}
