package topic21

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var p, q *ListNode = l1, l2
	var head *ListNode = &ListNode{}
	var cur *ListNode = head

	for p != nil && q != nil {
		if p.Val <= q.Val {
			cur.Next = p
			p = p.Next
		} else {
			cur.Next = q
			q = q.Next
		}
		cur = cur.Next
	}

	for p != nil {
		cur.Next = p
		p = p.Next
		cur = cur.Next
	}

	for q != nil {
		cur.Next = q
		q = q.Next
		cur = cur.Next
	}

	return head.Next
}
