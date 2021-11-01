package topic142

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var p, q *ListNode = head, head
	var step int = 0
	for {
		if q.Next != nil && q.Next.Next != nil {
			p = p.Next
			q = q.Next.Next
			step += 1
		} else {
			return nil
		}

		if p == q {
			break
		}
	}

	q = head
	for q != nil && p != nil {
		if q == p {
			return q
		}
		q = q.Next
		p = p.Next
	}

	return nil
}
