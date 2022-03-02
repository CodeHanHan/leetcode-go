package goden0204

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	var meta *ListNode = &ListNode{}
	meta.Next = head

	tmp := make([]*ListNode, 0)
	pre := meta
	p := head
	for p != nil {
		if p.Val < x {
			tmp = append(tmp, p)
			pre.Next = p.Next
		}
		pre = p
		p = p.Next
	}

	pre = meta
	p = pre.Next
	for p.Val < x {
		pre = p
		p = p.Next
	}

	for i := 0; i < len(tmp)-1; i++ {
		tmp[i].Next = tmp[i+1]
	}

	pre.Next = tmp[0]
	tmp[len(tmp)-1].Next = p

	return meta.Next
}
