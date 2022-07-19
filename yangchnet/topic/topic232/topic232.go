package topic232

type MyQueue struct {
	in  []int
	out []int
}

func Constructor() MyQueue {
	return MyQueue{
		in:  make([]int, 0),
		out: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

func (this *MyQueue) Pop() int {
	if this.Empty() {
		return -1
	}

	if len(this.out) <= 0 {
		this.out, this.in = this.in, this.out
	}

	ret := this.out[0]
	this.out = this.out[1:]

	return ret
}

func (this *MyQueue) Peek() int {
	if this.Empty() {
		return -1
	}

	if len(this.out) <= 0 {
		this.out, this.in = this.in, this.out
	}

	return this.out[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.in)+len(this.out) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
