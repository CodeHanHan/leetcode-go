package topic206

import (
	. "github.com/CodeHanHan/leetcode-go/base/LinkList"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	_head := &ListNode{Next: head}
	var fn func(head *ListNode) *ListNode
	fn = func(head *ListNode) *ListNode {
		if head.Next == nil {
			_head.Next = head
			return head
		}

		fn(head.Next).Next = head
		return head
	}

	fn(head).Next = nil

	return _head.Next
}

func reverseList1(head *ListNode) *ListNode {
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
