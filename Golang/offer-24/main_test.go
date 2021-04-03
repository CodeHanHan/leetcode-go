package main

import (
	"testing"
)

func Test_reverseList(t *testing.T) {
	head := &ListNode{
		Val:  1,
		Next: nil,
	}
	//var head *ListNode
	insertToTail(head, 2)
	insertToTail(head, 3)
	insertToTail(head, 4)
	insertToTail(head, 5)

	reverseList(head)
	//fmt.Printf("%v", h)

}
