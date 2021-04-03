package base

import (
	"bytes"
	"fmt"
)

/**
@Author: lichang
@Date: 3/29/2021
@Version: 1.0
@Des: 带有头节点的单链表
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildList(data []int) *ListNode {
	head := &ListNode{
		Val:  -1,
		Next: nil,
	}
	p := head
	for _, v := range data {
		p.Next = &ListNode{
			Val:  v,
			Next: nil,
		}
		p = p.Next
	}
	return head
}

func (l *ListNode) String() string {
	ret := bytes.Buffer{}
	p := l.Next
	for p != nil {
		ret.WriteString(fmt.Sprintf("%v → ", p.Val))
		p = p.Next
	}
	ret.WriteString("nil")
	return ret.String()
}
