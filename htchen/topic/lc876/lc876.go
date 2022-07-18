package lc876

import (
	. "github.com/CodeHanHan/leetcode-go/base/LinkList"
)

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
