package topic234

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

func isPalindrome(head *ListNode) bool {
	p := head

	ans := true
	var fn func(head *ListNode)
	fn = func(head *ListNode) {
		if head == nil {
			return
		}

		fn(head.Next)

		if p.Val != head.Val {
			ans = false
		}

		p = p.Next
	}

	fn(head)

	return ans
}
