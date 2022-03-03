package lc092

func reverseBetween1(head *ListNode, left int, right int) *ListNode {
	prehead := &ListNode{Next: head}
	pre := prehead
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	rightNode := pre
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}

	lefeNode := pre.Next
	after := rightNode.Next
	pre.Next = nil
	rightNode.Next = nil
	reverseList(lefeNode)
	pre.Next = rightNode
	lefeNode.Next = after

	return prehead.Next
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func reverseBetween2(head *ListNode, left int, right int) *ListNode {
	prehead := &ListNode{Next: head}
	pre := prehead

	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next

	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return prehead.Next
}
