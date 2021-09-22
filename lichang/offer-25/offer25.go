package offer25

import (
	"bytes"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{
		Val:  -1,
		Next: nil,
	}
	cur := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			cur = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			cur = l2
			l2 = l2.Next
		}
	}

	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}

	return head.Next
}

func (l *ListNode) String() string {
	ret := bytes.Buffer{}
	var p *ListNode = l
	if l.Val < 0 {
		p = l.Next
	}
	for p != nil {
		ret.WriteString(fmt.Sprintf("%v → ", p.Val))
		p = p.Next
	}
	ret.WriteString("nil")
	return ret.String()
}
