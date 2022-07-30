package topic146

import (
	"container/list"
)

type LRUCache struct {
	ll               *list.List
	items            map[int]*list.Element
	length, capacity int
}

func Constructor(capacity int) LRUCache {
	ll := list.New()

	return LRUCache{
		ll:       ll,
		items:    make(map[int]*list.Element),
		length:   0,
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.items[key]; !ok {
		return -1
	} else {
		this.ll.MoveToFront(node)
		return node.Value.(int)
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.items[key]; ok {
		node.Value = value
		this.ll.MoveToFront(node)
		return
	}

	if this.length >= this.capacity {
		delete(this.items, this.ll.Back().Value.(int))
		this.ll.Remove(this.ll.Back())
		this.length--
	}

	this.ll.PushFront(value)
	this.items[key] = this.ll.Front()
	this.length++
}
