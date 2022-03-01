package lc083

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	p := head

	for p != nil && p.Next != nil {
		if p.Val == p.Next.Val {
			x := p.Val
			for p.Next != nil && p.Next.Val == x {
				p.Next = p.Next.Next
			}
		} else {
			p = p.Next
		}
	}
	return dummy.Next
}
