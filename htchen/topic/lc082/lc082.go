package lc082

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	p := pre.Next
	for p != nil && p.Next != nil {
		if p.Val == p.Next.Val {
			x := p.Val
			for p != nil && p.Val == x {
				pre.Next = p.Next
				p = p.Next
			}
		} else {
			pre = pre.Next
			p = pre.Next
		}
	}
	return dummy.Next
}
