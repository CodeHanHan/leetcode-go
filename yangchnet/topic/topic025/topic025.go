package topic025

import "math"

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	var head_ *ListNode = &ListNode{math.MaxInt64, head}
	pre := head_

	var left, right *ListNode = head, head

	var i int = 1
	for right != nil && right.Next != nil {
		right = right.Next
		i++

		if i == k { // 此时从h到t为一段要反转的链表，暂时分离
			next := right.Next
			right.Next = nil
			h, t := reverseList(left)

			t.Next = next
			pre.Next = h

			pre = t
			left = pre.Next
			right = left
			i = 1
		}
	}

	return head_.Next
}

// 反转以head为头的链表
func reverseList(head *ListNode) (*ListNode, *ListNode) {
	if head == nil {
		return head, nil
	}
	var cur *ListNode
	var pre *ListNode
	cur = head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre, head
}
