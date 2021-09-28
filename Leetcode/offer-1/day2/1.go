package day2

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

func reversePrint(head *ListNode) (res []int) {
	if head == nil {
		return res
	}
	res = append(res, head.Val)
	res = append(reversePrint(head.Next), res...)
	return res

}

func reversePrintV2(head *ListNode) (res []int) {
	var list []int
	node := head
	for node != nil {
		list = append(list, node.Val)
		node = node.Next
	}
	for i := len(list) - 1; i >= 0; i-- {
		res = append(res, list[i])
	}
	return res
}

func reversePrintV3(head *ListNode) (res []int) {
	node := head
	for node != nil {
		res = append([]int{node.Val}, res...)
		node = node.Next
	}
	return res
}

func ReversePrint(head *ListNode) []int {
	return reversePrint(head)
}

func ReversePrintV2(head *ListNode) []int {
	return reversePrintV2(head)
}

func ReversePrintV3(head *ListNode) []int {
	return reversePrintV3(head)
}
