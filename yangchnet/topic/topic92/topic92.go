package topic92

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
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	_head := &ListNode{Val: math.MaxInt32, Next: head}
	var leftPtr, leftPre, rightPost, rightPtr *ListNode = nil, nil, nil, nil

	p := _head

	count := 0
	for count <= right {
		if count == left-1 {
			leftPtr = p.Next
			leftPre = p
		}
		if count == right {
			rightPtr = p
			rightPost = p.Next
		}
		p = p.Next
		count++
	}
	if leftPtr == nil || rightPtr == nil {
		return _head.Next
	}
	rightOrigin := rightPtr

	for leftPre.Next != rightOrigin {
		// 重新链接链表
		rightPtr.Next = leftPtr
		leftPre.Next = leftPtr.Next
		leftPtr.Next = rightPost

		// 重新赋值指针
		leftPtr = leftPre.Next
		rightPost = rightPtr.Next
		// rightPtr = rightPtr.Next
	}

	return _head.Next
}
