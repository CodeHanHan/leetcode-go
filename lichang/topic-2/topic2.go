package topic2

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int = 0

	var p1, p2 *ListNode = l1, l2
	var cur *ListNode = &ListNode{Val: -1, Next: nil}
	var head *ListNode = cur

	for p1 != nil && p2 != nil {
		newNode := &ListNode{
			Val:  (p1.Val + p2.Val + carry) % 10,
			Next: nil,
		}
		carry = (p1.Val + p2.Val) / 10
		cur.Next = newNode
		cur = newNode

		p1 = p1.Next
		p2 = p2.Next
	}

	for p1 != nil {
		newNode := &ListNode{
			Val:  (p1.Val + carry) % 10,
			Next: nil,
		}
		cur.Next = newNode
		cur = newNode
		carry = (p1.Val + carry) / 10

		p1 = p1.Next
	}

	for p2 != nil {
		newNode := &ListNode{
			Val:  (p2.Val + carry) % 10,
			Next: nil,
		}
		cur.Next = newNode
		cur = newNode
		carry = (p2.Val + carry) / 10

		p2 = p2.Next
	}

	if carry != 0 {
		cur.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}

	return head.Next
}

// 内存优化
func addTwoNumbers_1(l1 *ListNode, l2 *ListNode) *ListNode {
	var p1, p2, cur *ListNode
	var carry, c int

	p1, p2 = l1, l2

	for p1 != nil && p2 != nil {
		cur = p1

		c = carry

		carry = (p1.Val + p2.Val + c) / 10
		p1.Val = (p1.Val + p2.Val + c) % 10

		p1 = p1.Next
		p2 = p2.Next
	}

	for p1 != nil {
		cur = p1
		c = carry
		carry = (p1.Val + c) / 10
		p1.Val = (p1.Val + c) % 10

		p1 = p1.Next
	}

	if p2 != nil {
		cur.Next = p2
	}

	for p2 != nil {
		cur = p2
		c = carry

		carry = (p2.Val + c) / 10
		p2.Val = (p2.Val + c) % 10

		p2 = p2.Next
	}

	if carry != 0 {
		cur.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}
	return l1
}
