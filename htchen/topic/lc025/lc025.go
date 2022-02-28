package lc025

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{0, head}
	pre := dummy
	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummy.Next
			}
		}
		next := tail.Next
		head, tail = reverseList(head, tail)
		pre.Next = head
		tail.Next = next
		pre = tail
		head = next
	}

	return dummy.Next
}

func reverseList(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	var prev *ListNode
	prev = tail.Next
	p := head
	for prev != tail {
		next := p.Next
		p.Next = prev
		prev = p
		p = next
	}
	
	return tail, head
}
