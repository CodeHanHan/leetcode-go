package offer_09_s

type CQueue struct {
	head []int
	tail []int
}

func Constructor() CQueue {
	return CQueue{
		head: make([]int, 0), // 0为栈底，最后一个元素为top
		tail: make([]int, 0), // 0为栈底，最后一个元素为top
	}
}

func (this *CQueue) AppendTail(value int) {
	this.tail = append(this.tail, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.head) > 0 {
		val := this.head[len(this.head)-1]
		this.head = this.head[:len(this.head)-1]
		return val
	} else {
		if len(this.tail) > 0 {
			for len(this.tail) > 0 {
				last := this.tail[len(this.tail)-1]
				this.tail = this.tail[:len(this.tail)-1]
				this.head = append(this.head, last)
			}
			val := this.head[len(this.head)-1]
			this.head = this.head[:len(this.head)-1]
			return val
		} else {
			return -1
		}
	}
}
