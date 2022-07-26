package offer22

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
func getKthFromEnd(head *ListNode, k int) *ListNode {
	var ret *ListNode

	_k := 0
	var fn func(head *ListNode)
	fn = func(head *ListNode) {
		if head == nil {
			return
		}

		fn(head.Next)
		_k++
		if _k == k {
			ret = head
			_k = k + 1
		}
	}

	fn(head)

	return ret
}
