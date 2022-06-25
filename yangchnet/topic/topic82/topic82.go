package topic82

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
func deleteDuplicates(head *ListNode) *ListNode {
	_head := &ListNode{Next: head}
	pre, slow, fast := _head, head, head

	for fast != nil {
		for fast.Next != nil && fast.Next.Val == slow.Val {
			fast = fast.Next
		}

		if fast != slow {
			pre.Next = fast.Next
		} else {
			pre = fast
		}

		slow = fast.Next
		fast = slow
	}

	return _head.Next
}
