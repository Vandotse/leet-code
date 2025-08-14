/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	groupPrev := dummy
	for true {
		kth := find(groupPrev, k)
		if kth == nil {
			break
		}
		groupNext := kth.Next

		prev := kth.Next
		curr := groupPrev.Next
		for curr != groupNext {
			temp := curr.Next
			curr.Next = prev
			prev = curr
			curr = temp
		}

		temp := groupPrev.Next
		groupPrev.Next = kth
		groupPrev = temp
	}
	return dummy.Next

}

func find(curr *ListNode, k int) *ListNode {
	for i := 0; i < k && curr != nil; i++ {
		curr = curr.Next
	}
	return curr
}