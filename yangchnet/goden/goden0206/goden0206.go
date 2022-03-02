package goden0206

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	ch := make(chan bool, 1)
	p := head
	var fn func(head *ListNode)
	fn = func(head *ListNode) {
		if head == nil {
			return
		}

		fn(head.Next)

		if p.Val != head.Val && len(ch) < 1 {
			ch <- false
			return
		}

		if p == head && len(ch) < 1 {
			ch <- true
			return
		}

		p = p.Next
	}

	fn(head)

	if len(ch) < 1 {
		return true
	}

	return <-ch
}
