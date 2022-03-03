package lc086

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

// Len return list length, when list has cycle structure, return -1
func (l *ListNode) Len() int {
	if l == nil {
		return 0
	}

	if ok, _ := l.HasCycle(); ok {
		return -1
	}

	var p *ListNode = l
	length := 0
	for p != nil {
		length += 1
		p = p.Next
	}

	return length
}

// Tail return last list node
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

// HasCycle check if linklist has cycle structure, if not, return false, nil, if yes, return true and the cycle start node
func (l *ListNode) HasCycle() (bool, *ListNode) {
	if l == nil {
		return false, nil
	}

	var p, q *ListNode = l, l
	var step int = 0
	for {
		if q.Next != nil && q.Next.Next != nil {
			p = p.Next
			q = q.Next.Next
			step += 1
		} else {
			return false, nil
		}

		if p == q {
			break
		}
	}

	var r *ListNode = l
	for r != nil && p != nil {
		if r == p {
			return true, r
		}
		r = r.Next
		p = p.Next
	}

	return false, nil
}

// Index return the node which value equals target
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
	l := BuildListWithNoHead(list)
	tail := l.Tail()
	p := l.Index(list[cursor])
	tail.Next = p
	return l
}

// Slice return slice which contains all node's values in order.
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

// String implements Stringer interface.
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
