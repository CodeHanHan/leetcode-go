package goden1625

import (
	"container/list"
)

type LRUCache struct {
	ll  *list.List
	m   map[int]*list.Element
	cap int
}

type KV struct {
	K int
	V int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		ll:  list.New(),
		m:   make(map[int]*list.Element),
		cap: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	ret, ok := this.m[key]
	if !ok {
		return -1
	}

	this.ll.MoveToFront(ret)

	return ret.Value.(*KV).V
}

func (this *LRUCache) Put(key int, value int) {
	if e, ok := this.m[key]; ok {
		e.Value.(*KV).V = value
		this.ll.MoveToFront(e)
		return
	}

	if this.ll.Len() >= this.cap {
		delete(this.m, this.ll.Back().Value.(*KV).K)
		this.ll.Remove(this.ll.Back())
	}

	e := this.ll.PushFront(&KV{
		K: key,
		V: value,
	})
	this.m[key] = e
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
