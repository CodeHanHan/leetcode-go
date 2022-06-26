package topic83

import (
	"math"

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
	_head := &ListNode{
		Val:  math.MaxInt32,
		Next: head,
	}

	pre, slow, fast := _head, head, head

	for fast != nil {
		for fast.Next != nil && fast.Next.Val == slow.Val {
			fast = fast.Next
		}

		if slow != fast {
			pre.Next = fast
		}

		pre = fast
		slow = fast.Next
		fast = slow
	}

	return _head.Next
}
