package lc160

import (
	. "github.com/CodeHanHan/leetcode-go/base/LinkList"
)

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	visited := map[*ListNode]bool{}
	for node := headA; node != nil; node = node.Next {
		visited[node] = true
	}
	for node := headB; node != nil; node = node.Next {
		if visited[node] {
			return node
		}
	}
	return nil
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	lenA, lenB := lengthOfList(headA), lengthOfList(headB)
	if lenA > lenB {
		len := lenA - lenB
		for i := 0; i < len; i++ {
			headA = headA.Next
		}
	} else {
		len := lenB - lenA
		for i := 0; i < len; i++ {
			headB = headB.Next
		}
	}

	for headA != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}

func lengthOfList(head *ListNode) int {
	if head == nil {
		return 0
	}
	p := head
	length := 0
	for p != nil {
		length++
		p = p.Next
	}
	return length
}
