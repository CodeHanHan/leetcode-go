package main

type ListNode struct {
	Val  int
	Next *ListNode
}

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
