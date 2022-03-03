package lc086

func partition(head *ListNode, x int) *ListNode {
	smalllist := &ListNode{Next: head}
	small := smalllist
	largelist := &ListNode{Next: head}
	large := largelist
	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}
		head = head.Next
	}
	
	large.Next = nil
	small.Next = largelist.Next

	return smalllist.Next
}
