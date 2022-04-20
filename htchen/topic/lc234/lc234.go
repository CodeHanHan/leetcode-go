package lc234

import (
	. "github.com/CodeHanHan/leetcode-go/base/LinkList"
)

func isPalindrome1(head *ListNode) bool {
	vals := []int{}
	for ; head != nil; head = head.Next {
		vals = append(vals, head.Val)
	}
	n := len(vals)
	for i, v := range vals[:n/2] {
		if v != vals[n-i-1] {
			return false
		}
	}
	return true
}

func isPalindrome2(head *ListNode) bool {
	if head == nil {
		return false
	}

	firstHalfEnd := endOfFirstHalf(head)
	secondHalfStart := reverseList(firstHalfEnd.Next)

	node1 := head
	node2 := secondHalfStart
	for node2 != nil {
		if node1.Val != node2.Val {
			return false
		}
		node1 = node1.Next
		node2 = node2.Next
	}
    firstHalfEnd.Next = reverseList(secondHalfStart)
	return true
}

func endOfFirstHalf(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
