package offer_06_s

import "fmt"

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

func reversePrint(head *ListNode) []int {
	ret := make([]int, 0)
	var f func(p *ListNode)
	f = func(p *ListNode) {
		if p == nil {
			return
		}
		f(p.Next)
		ret = append(ret, p.Val)
	}
	f(head.Next) // 若链表无头，则参数为head
	fmt.Println(ret)
	return ret
}
