/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	front := head
	back := dummy
	for n > 0 && front != nil {
		front = front.Next
		n--
	}
	for front != nil {
		front = front.Next
		back = back.Next
	}

	back.Next = back.Next.Next
	return dummy.Next
}