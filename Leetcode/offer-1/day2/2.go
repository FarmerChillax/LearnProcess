package day2

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) (res *ListNode) {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	rear := reverseList(head.Next)
	head.Next.Next = head
	return rear
}
