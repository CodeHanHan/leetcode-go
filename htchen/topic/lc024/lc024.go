package lc024

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	temp := dummy

	for temp.Next != nil && temp.Next.Next != nil {
		p1, p2 := temp.Next, temp.Next.Next
		temp.Next = p2
		p1.Next = p2.Next
		p2.Next = p1
		temp = p1
	}
	return dummy.Next
}
