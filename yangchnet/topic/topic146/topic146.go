package topic146

import (
	"math"
)

type ListNode struct {
	Key  int
	Val  int
	Next *ListNode
	Pre  *ListNode
}

type LRUCache struct {
	head     *ListNode
	tail     *ListNode
	items    map[int]*ListNode
	length   int
	capacity int
}

func Constructor(capacity int) LRUCache {
	head := &ListNode{
		Key:  math.MinInt32,
		Val:  math.MinInt32,
		Next: nil,
		Pre:  nil,
	}

	tail := &ListNode{
		Key:  math.MaxInt32,
		Val:  math.MaxInt32,
		Next: nil,
		Pre:  nil,
	}

	head.Next = tail
	tail.Pre = head

	return LRUCache{
		head:     head,
		tail:     tail,
		items:    make(map[int]*ListNode),
		length:   0,
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.items[key]; ok {
		node.Pre.Next = node.Next
		node.Next.Pre = node.Pre
		this.head.HeadInsert(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// 可以直接找到
	if node, ok := this.items[key]; ok {
		node.Val = value

		node.Pre.Next = node.Next
		node.Next.Pre = node.Pre
		this.head.HeadInsert(node)
		return
	}

	// 链未满
	if this.length < this.capacity {
		node := this.head.HeadInsert(newNode(key, value))
		this.length++
		this.items[key] = node
	} else { // 链已满
		abd := this.tail.Pre
		delete(this.items, abd.Key)

		abd.Pre.Next = abd.Next
		abd.Next.Pre = abd.Pre

		node := this.head.HeadInsert(newNode(key, value))
		this.items[key] = node
	}
}

func newNode(key, value int) *ListNode {
	return &ListNode{
		Key: key,
		Val: value,
	}
}

func (head *ListNode) HeadInsert(node *ListNode) *ListNode {
	node.Next = head.Next
	node.Pre = head
	head.Next.Pre = node
	head.Next = node
	return node
}
