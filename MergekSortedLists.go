/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	for len(lists) > 1 {
		var merged []*ListNode
		for i := 0; i < len(lists); i += 2 {
			l1 := lists[i]
			var l2 *ListNode
			if i+1 < len(lists) {
				l2 = lists[i+1]
			}
			merged = append(merged, mergeList(l1, l2))
		}
		lists = merged
	}
	return lists[0]
}

func mergeList(first *ListNode, second *ListNode) *ListNode {
	dummy := &ListNode{}
	op := dummy
	for first != nil && second != nil {
		if first.Val <= second.Val {
			op.Next = first
			first = first.Next
		} else {
			op.Next = second
			second = second.Next
		}
		op = op.Next
	}
	if first != nil {
		op.Next = first
	}
	if second != nil {
		op.Next = second
	}
	return dummy.Next
}