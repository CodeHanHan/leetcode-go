package lc146

type DlinkNode struct {
	key, value int
	next, prev *DlinkNode
}

type LRUCache struct {
	capacity   int
	size       int
	cache      map[int]*DlinkNode
	head, tail *DlinkNode
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		size: 0,
		cache:    map[int]*DlinkNode{},
		head:     initDlinkNode(0, 0),
		tail:     initDlinkNode(0, 0),
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		node := initDlinkNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			removedNode := this.removeTail()
			delete(this.cache, removedNode.key)
			this.size--
		}
	} else {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}

func initDlinkNode(key, value int) *DlinkNode {
	return &DlinkNode{
		key:   key,
		value: value,
	}
}

func (this *LRUCache) addToHead(node *DlinkNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DlinkNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DlinkNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DlinkNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}
