package topic61

import "math"

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil { // head为空，直接返回
		return nil
	}

	var tail *ListNode = head
	var length int = 1     // 链表长度
	for tail.Next != nil { // 将指针调整到最后一个元素
		tail = tail.Next
		length++
	}

	if length == 1 { // 链表只有一个节点，返回其自身
		return head
	}

	l := length - (k % length) // 将链表向右旋转k个位置，相当于将链表向左旋转length-k个位置，又由于k可能大于length, 因此取余

	var head_ *ListNode = &ListNode{
		Val:  math.MaxInt32,
		Next: nil,
	}
	head_.Next = head // 增加一个头节点，便于操作

	for i := 0; i < l; i++ {
		tmp := head_.Next.Next
		tail.Next = head_.Next
		head_.Next = tmp
		tail = tail.Next
		tail.Next = nil
	}

	return head_.Next
}
