package lc143

import (
	. "github.com/CodeHanHan/leetcode-go/base/LinkList"
)

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	nodes := []*ListNode{}
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	low, high := 0, len(nodes)-1
	for low < high {
		nodes[low].Next = nodes[high]
		low++
		if low == high {
			break
		}
		nodes[high].Next = nodes[low]
		high--
	}
	nodes[low].Next = nil
}

func reorderList1(head *ListNode) {
	if head == nil {
		return
	}
	mid := middleNode(head)
	l1 := head
	l2 := mid.Next
	mid.Next = nil
	l2 = reverseList(l2)
	mergeList(l1, l2)
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	p := head
	for p != nil {
		nex := p.Next
		p.Next = prev
		prev = p
		p = nex
	}
	return prev
}

func mergeList(l1, l2 *ListNode) {
	l1tmp, l2tmp := l1, l2
	for l1 != nil && l2 != nil {
		l1tmp = l1.Next
		l2tmp = l2.Next

		l1.Next = l2
		l1 = l1tmp

		l2.Next = l1
		l2 = l2tmp
	}
}
