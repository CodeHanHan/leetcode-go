package topic148

import "sort"

// O(nlogn)
func sortList_1(head *ListNode) *ListNode {
	nums := make([]int, 0)
	var p *ListNode = head

	for p != nil {
		nums = append(nums, p.Val)
		p = p.Next
	}

	sort.Ints(nums)

	p = head
	for i := 0; i < len(nums); i++ {
		p.Val = nums[i]
		p = p.Next
	}

	return head
}

// 归并法
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	l1, l2 := splitList(head)

	return mergeTwoLists(sortList(l1), sortList(l2))
}

func splitList(head *ListNode) (l1 *ListNode, l2 *ListNode) {
	if head == nil {
		return nil, nil
	}

	var slow, fast *ListNode = head, head

	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	l1 = head
	l2 = slow.Next
	slow.Next = nil
	return l1, l2
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var p, q *ListNode = l1, l2
	var head *ListNode = &ListNode{}
	var cur *ListNode = head

	for p != nil && q != nil {
		if p.Val <= q.Val {
			cur.Next = p
			p = p.Next
		} else {
			cur.Next = q
			q = q.Next
		}
		cur = cur.Next
	}

	for p != nil {
		cur.Next = p
		p = p.Next
		cur = cur.Next
	}

	for q != nil {
		cur.Next = q
		q = q.Next
		cur = cur.Next
	}

	return head.Next
}
