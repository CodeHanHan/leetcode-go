package offer24

import (
	. "github.com/CodeHanHan/leetcode-go/base/LinkList"
)

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
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
	return pre
}

func insertToTail(head *ListNode, val int) *ListNode {
	var cur *ListNode = head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &ListNode{
		Val:  val,
		Next: nil,
	}
	return head
}

func traverseList(head *ListNode) []int {
	nums := make([]int, 0)

	var p *ListNode = head
	for p != nil {
		nums = append(nums, p.Val)
		p = p.Next
	}
	return nums
}

func reverseList2(head *ListNode) *ListNode {
	var res *ListNode
	var fn func(pre, cur *ListNode, op func(pre, cur *ListNode))
	fn = func(pre, cur *ListNode, op func(pre, cur *ListNode)) {
		if cur == nil {
			res = pre
			return
		}

		fn(cur, cur.Next, op)
		op(pre, cur)
	}

	fn(nil, head, func(pre, cur *ListNode) {
		cur.Next = pre
	})

	return res
}
