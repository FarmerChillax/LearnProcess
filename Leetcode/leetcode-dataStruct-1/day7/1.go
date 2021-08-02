package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//  my hashMap func
// func hasCycle(head *ListNode) bool {
// 	nodeMap := map[ListNode]int{}
// 	node := head
// 	for {
// 		if node == nil {
// 			return false
// 		}
// 		nodeMap[*node]++
// 		if nodeMap[*node] > 1 {
// 			return true
// 		}
// 		node = node.Next
// 	}
// }

// leetcode hashMap
func hasCycle(head *ListNode) bool {
	seen := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := seen[head]; ok {
			return true
		}
		seen[head] = struct{}{}
		head = head.Next
	}
	return false
}

func main() {
	inputList := []int{3, 2, 0, -4}
	pos := 1
	listPoint := newListNode(inputList, pos)
	// listPoint.PrintNode(2)
	// fmt.Printf("%+v, %T\n", listPoint.Next, listPoint.Next)
	fmt.Println(hasCycle(listPoint))
}

//
type ListNode struct {
	Val  int
	Next *ListNode
}

func newListNode(inputList []int, pos int) *ListNode {
	var target *ListNode = nil
	head := &ListNode{}
	node := head

	for index, input := range inputList {
		// fmt.Println(input)
		tmp := &ListNode{
			Val:  input,
			Next: nil,
		}
		if index == pos {
			target = tmp
		}
		node.Next = tmp
		// fmt.Println(node)
		node = node.Next
	}
	node.Next = target
	return head.Next
}

func (n *ListNode) PrintNode(posValue int) {
	node := n
	flag := 0
	for {
		if node == nil || flag == 2 {
			return
		}
		if node.Val == posValue {
			flag++
		}

		fmt.Println(node.Val)
		node = node.Next
	}
}
