package topic143

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
func reorderList(head *ListNode) {
	queue := make([]*ListNode, 0)

	p := head
	for p != nil {
		queue = append(queue, p)
		p = p.Next
	}

	p, q := head, queue[len(queue)-1]
	for p != q && p.Next != q {
		qPre := queue[len(queue)-2]
		qPre.Next = q.Next

		q.Next = p.Next
		p.Next = q

		p = p.Next.Next
		queue = queue[:len(queue)-1]
		q = queue[len(queue)-1]
	}

}
