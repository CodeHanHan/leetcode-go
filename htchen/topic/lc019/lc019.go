package lc019

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pre := &ListNode{0, head}
	fast, slow := pre, pre
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return pre.Next
}
