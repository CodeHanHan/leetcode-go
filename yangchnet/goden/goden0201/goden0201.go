package goden0201

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	existed := make(map[int]bool)
	p := head
	existed[p.Val] = true
	q := head.Next

	for q != nil {
		if !existed[q.Val] {
			p.Next.Val = q.Val
			existed[q.Val] = true
			p = p.Next
		}
		q = q.Next
	}

	p.Next = nil

	return head
}
