package goden0302

import "math"

type MinStack struct {
	nums []int // nums[0]为哨兵，初始为-1
	min  []int // 与nums等长，top指向的位置为栈中最小值, min[0] 为哨兵，值为math.Max
	top  int   // 初始为0,指向当前栈顶元素, 为0时为空，为len(nums)-1时为满
}

/** initialize your data structure here. */
func Constructor() MinStack {
	s := MinStack{
		nums: make([]int, 100),
		min:  make([]int, 100),
		top:  0,
	}

	s.nums[0] = -1
	s.min[0] = math.MaxInt64
	return s
}

func (this *MinStack) Push(x int) {
	if this.top >= len(this.nums)-1 { // 栈满
		this.nums = append(this.nums, make([]int, len(this.nums))...)
		this.min = append(this.min, make([]int, len(this.min))...)
	}

	min := this.min[this.top]
	this.top++
	this.nums[this.top] = x
	if x < min {
		min = x
	}
	this.min[this.top] = min
}

func (this *MinStack) Pop() {
	if this.IsEmpty() {
		return
	}

	this.top--
}

func (this *MinStack) Top() int {
	return this.nums[this.top]
}

func (this *MinStack) GetMin() int {
	if this.IsEmpty() {
		return -1
	}
	return this.min[this.top]
}

func (this *MinStack) IsEmpty() bool {
	return this.top == 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
