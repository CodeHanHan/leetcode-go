package topic146

// type ListNode struct {
// 	Key  int
// 	Val  int
// 	Next *ListNode
// 	Pre  *ListNode
// }

// type LRUCache struct {
// 	head     *ListNode
// 	tail     *ListNode
// 	items    map[int]*ListNode
// 	length   int
// 	capacity int
// }

// func Constructor(capacity int) LRUCache {
// 	head := &ListNode{
// 		Key:  math.MinInt32,
// 		Val:  math.MinInt32,
// 		Next: nil,
// 		Pre:  nil,
// 	}

// 	tail := &ListNode{
// 		Key:  math.MaxInt32,
// 		Val:  math.MaxInt32,
// 		Next: nil,
// 		Pre:  nil,
// 	}

// 	head.Next = tail
// 	tail.Pre = head

// 	return LRUCache{
// 		head:     head,
// 		tail:     tail,
// 		items:    make(map[int]*ListNode),
// 		length:   0,
// 		capacity: capacity,
// 	}
// }

// func (this *LRUCache) Get(key int) int {
// 	if node, ok := this.items[key]; ok {
// 		node.Pre.Next = node.Next
// 		node.Next.Pre = node.Pre
// 		this.head.HeadInsert(node)
// 		return node.Val
// 	}
// 	return -1
// }

// func (this *LRUCache) Put(key int, value int) {
// 	// 可以直接找到
// 	if node, ok := this.items[key]; ok {
// 		node.Val = value

// 		node.Pre.Next = node.Next
// 		node.Next.Pre = node.Pre
// 		this.head.HeadInsert(node)
// 		return
// 	}

// 	// 链未满
// 	if this.length < this.capacity {
// 		node := this.head.HeadInsert(newNode(key, value))
// 		this.length++
// 		this.items[key] = node
// 	} else { // 链已满
// 		abd := this.tail.Pre
// 		delete(this.items, abd.Key)

// 		abd.Pre.Next = abd.Next
// 		abd.Next.Pre = abd.Pre

// 		node := this.head.HeadInsert(newNode(key, value))
// 		this.items[key] = node
// 	}
// }

// func newNode(key, value int) *ListNode {
// 	return &ListNode{
// 		Key: key,
// 		Val: value,
// 	}
// }

// func (head *ListNode) HeadInsert(node *ListNode) *ListNode {
// 	node.Next = head.Next
// 	node.Pre = head
// 	head.Next.Pre = node
// 	head.Next = node
// 	return node
// }

// --------------------------------------------

// type LRUCache struct {
// 	head, tail       *ListNode
// 	items            map[int]*ListNode
// 	length, capacity int
// }

// type ListNode struct {
// 	Key  int
// 	Val  int
// 	Next *ListNode
// 	Pre  *ListNode
// }

// func Constructor(capacity int) LRUCache {
// 	head := &ListNode{
// 		Key: math.MinInt32,
// 		Val: math.MinInt32,
// 		Pre: nil,
// 	}

// 	tail := &ListNode{
// 		Key:  math.MaxInt32,
// 		Val:  math.MaxInt32,
// 		Next: nil,
// 	}

// 	head.Next = tail
// 	tail.Pre = head

// 	return LRUCache{
// 		head:     head,
// 		tail:     tail,
// 		items:    make(map[int]*ListNode),
// 		length:   0,
// 		capacity: capacity,
// 	}
// }

// func (this *LRUCache) Get(key int) int {
// 	if node, ok := this.items[key]; !ok {
// 		return -1
// 	} else {
// 		HeadInsert(this.head, CutNode(node))
// 		return node.Val
// 	}
// }

// func (this *LRUCache) Put(key int, value int) {
// 	// key 已存在
// 	if node, ok := this.items[key]; ok {
// 		node.Val = value
// 		HeadInsert(this.head, CutNode(node))
// 		return
// 	}

// 	if this.length >= this.capacity {
// 		delete(this.items, CutNode(this.tail.Pre).Val)
// 		this.length--
// 	}

// 	newNode := &ListNode{
// 		Key: key,
// 		Val: value,
// 	}

// 	HeadInsert(this.head, newNode)
// 	this.items[key] = newNode

// 	this.length++
// }

// func HeadInsert(head, node *ListNode) {
// 	node.Next = head.Next
// 	node.Pre = head

// 	head.Next = node
// 	node.Next.Pre = node
// }

// func CutNode(node *ListNode) *ListNode {
// 	node.Pre.Next = node.Next
// 	node.Next.Pre = node.Pre

// 	return node
// }
