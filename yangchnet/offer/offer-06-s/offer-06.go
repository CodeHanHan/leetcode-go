package offer_06_s

import (
	. "github.com/CodeHanHan/leetcode-go/base/LinkList"
)

func reversePrint(head *ListNode) []int {
	ret := make([]int, 0)

	var f func(p *ListNode, op func(node *ListNode))
	f = func(p *ListNode, op func(node *ListNode)) {
		if p == nil {
			return
		}
		f(p.Next, op)
		op(p)
	}

	f(head.Next, func(node *ListNode) {
		ret = append(ret, node.Val)
	})

	return ret
}
