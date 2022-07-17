package topic019

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	order := 0
	var pre *ListNode = &ListNode{}
	pre.Next = head

	var fn func(head *ListNode, pre *ListNode) int
	fn = func(head *ListNode, pre *ListNode) int {
		if head == nil {
			return -1
		}

		order += fn(head.Next, head)

		if order == -n { // head即为要删除的节点
			pre.Next = head.Next
		}

		return -1
	}

	fn(head, pre)

	return pre.Next
}

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	k := 1

	_head := &ListNode{Next: head}

	var fn func(node, pre *ListNode)
	fn = func(node, pre *ListNode) {
		if node == nil {
			return
		}

		fn(node.Next, node)

		if k == n {
			pre.Next = node.Next
			k = -1
			return
		}
		k++
	}

	fn(head, _head)

	return _head.Next
}
