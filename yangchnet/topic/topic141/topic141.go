package topic141

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	var p, q *ListNode = head, head

	for {
		if q.Next != nil && q.Next.Next != nil {
			p = p.Next
			q = q.Next.Next
		} else {
			return false
		}

		if p == q {
			return true
		}
	}
}
