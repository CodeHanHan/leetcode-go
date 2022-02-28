package goden0202

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func kthToLast_1(head *ListNode, k int) int {
	var p, q *ListNode = head, head
	for i := 0; i < k; i++ {
		q = q.Next
	}

	for q != nil {
		p = p.Next
		q = q.Next
	}

	return p.Val
}

func kthToLast(head *ListNode, k int) int {
	var c, val int
	var fn func(head *ListNode)
	fn = func(p *ListNode) {
		if p == nil {
			return
		}
		fn(p.Next)

		c++
		if c == k {
			val = p.Val
		}
	}

	fn(head)

	return val
}
