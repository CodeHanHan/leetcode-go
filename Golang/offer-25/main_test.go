package main

import "testing"

func Test_mergeTwoLists(t *testing.T) {
	head1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	//var head *ListNode
	insertToTail(head1, 3)
	insertToTail(head1, 4)
	insertToTail(head1, 6)
	insertToTail(head1, 8)
	head1.print()

	head2 := &ListNode{
		Val:  2,
		Next: nil,
	}
	//var head *ListNode
	insertToTail(head2, 3)
	insertToTail(head2, 5)
	insertToTail(head2, 7)
	insertToTail(head2, 9)
	head2.print()

	mergeTwoLists(head1, head2)
	head1.print()

}
