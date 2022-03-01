package lc061


func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}

	len := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		len++
	}
	if len == k {
		return head
	}
	tail.Next = head
	for i := 0; i < (len - (k % len)); i++ {
		tail = tail.Next
	}
	res := tail.Next
	tail.Next = nil

	return res
}
