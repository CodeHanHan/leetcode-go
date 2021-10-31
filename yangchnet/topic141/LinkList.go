package topic141

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

func BuildListWithHead(data []int) *ListNode {
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

func BuildListWithNoHead(nums []int) *ListNode {
	if len(nums) < 1 {
		return nil
	}

	head := &ListNode{
		Val:  nums[0],
		Next: nil,
	}

	for i := 1; i < len(nums); i++ {
		insertToTail(head, nums[i])
	}
	return head
}

func (l *ListNode) Tail() *ListNode {
	if l == nil {
		return nil
	}

	p := l
	for p.Next != nil {
		p = p.Next
	}

	return p
}

func (l *ListNode) Index(target int) *ListNode {
	if l == nil {
		return nil
	}

	p := l
	for p != nil {
		if p.Val == target {
			return p
		}
		p = p.Next
	}

	return nil
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

// BuildCircleListWithNoHead build a circle list, which tail point to list[cursor]
func BuildCircleListWithNoHead(list []int, cursor int) *ListNode {
	if cursor >= len(list) {
		return nil
	}

	l := BuildListWithNoHead(list)
	tail := l.Tail()
	p := l.Index(list[cursor])
	tail.Next = p
	return l
}

func (l *ListNode) Slice() []int {
	var p *ListNode = l
	if l.Val < 0 {
		p = l.Next
	}
	a := make([]int, 0)
	for p != nil {
		a = append(a, p.Val)
		p = p.Next
	}
	return a
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
