package goden0205

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var c, sum, s int
	var head *ListNode = &ListNode{}
	var p, q, v *ListNode = l1, l2, head
	for p != nil && q != nil {
		sum = p.Val + q.Val + c
		c = sum / 10
		s = sum % 10
		v.Next = &ListNode{Val: s}
		p = p.Next
		q = q.Next
		v = v.Next
	}

	for p != nil {
		sum = p.Val + c
		c = sum / 10
		s = sum % 10
		v.Next = &ListNode{Val: s}
		v = v.Next
		p = p.Next
	}

	for q != nil {
		sum = q.Val + c
		c = sum / 10
		s = sum % 10
		v.Next = &ListNode{Val: s}
		v = v.Next
		q = q.Next
	}

	if c > 0 {
		v.Next = &ListNode{Val: c}
	}

	return head.Next
}
