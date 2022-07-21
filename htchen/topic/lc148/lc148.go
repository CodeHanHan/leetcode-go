package lc148

import (
	. "github.com/CodeHanHan/leetcode-go/base/LinkList"
)

func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	mid := slow
	return merge(sort(head, mid), sort(mid, tail))
}

func merge(head1, head2 *ListNode) *ListNode {
	dummy := &ListNode{}
	temp, temp1, temp2 := dummy, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}
	return dummy.Next
}

func sortList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}

	dummyNode := &ListNode{Next: head}
	for mlen := 1; mlen < length; mlen*=2 {
		prev, cur := dummyNode, dummyNode.Next
        for cur!=nil{
            head1 := cur
            for i := 1; i < mlen && cur.Next != nil; i++ {
                cur = cur.Next
            }

            head2 := cur.Next
            cur.Next = nil
            cur = head2
            for i := 1; i < mlen && cur != nil && cur.Next != nil; i++ {
                cur = cur.Next
            }
            var next *ListNode
            if cur!=nil{
                next = cur.Next
                cur.Next=nil
            }
            
            prev.Next = merge(head1, head2)
            
            for prev.Next != nil {
                prev = prev.Next
            }
            cur = next
        }
	}
	return dummyNode.Next
}