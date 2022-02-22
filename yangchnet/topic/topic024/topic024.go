package topic024

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	picket := &ListNode{}
	picket.Next = head

	var pre, mid, post *ListNode = picket, picket.Next, picket.Next.Next
	for {
		pre.Next = post
		mid.Next = post.Next
		post.Next = mid

		if mid.Next != nil && mid.Next.Next != nil {
			pre = mid
			mid = pre.Next
			post = mid.Next
		} else {
			break
		}
	}

	return picket.Next
}
