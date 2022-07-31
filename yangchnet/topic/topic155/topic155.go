package topic155

import "math"

// type MinStack struct {
// 	nums []int
// 	ms   []int
// }

// func Constructor() MinStack {
// 	return MinStack{
// 		nums: make([]int, 0),
// 		ms:   []int{math.MaxInt32},
// 	}
// }

// func (this *MinStack) Push(val int) {
// 	this.nums = append(this.nums, val)
// 	if val < this.ms[len(this.ms)-1] {
// 		this.ms = append(this.ms, val)
// 	} else {
// 		this.ms = append(this.ms, this.ms[len(this.ms)-1])
// 	}
// }

// func (this *MinStack) Pop() {
// 	this.nums = this.nums[:len(this.nums)-1]
// 	this.ms = this.ms[:len(this.ms)-1]
// }

// func (this *MinStack) Top() int {
// 	return this.nums[len(this.nums)-1]
// }

// func (this *MinStack) GetMin() int {
// 	return this.ms[len(this.ms)-1]
// }

// ----------------------------

// 不使用额外空间
type MinStack struct {
	nums []int
	minV int
}

func Constructor() MinStack {
	return MinStack{
		nums: make([]int, 0),
		minV: math.MaxInt32,
	}
}

func (this *MinStack) Push(val int) {
	if len(this.nums) <= 0 {
		this.nums = append(this.nums, 0)
		this.minV = val
	} else {
		diff := val - this.minV
		this.nums = append(this.nums, diff)
		if diff <= 0 {
			this.minV = val
		}
	}
}

func (this *MinStack) Pop() {
	if len(this.nums) <= 0 {
		return
	}
	diff := this.nums[len(this.nums)-1]
	this.nums = this.nums[:len(this.nums)-1]
	if diff < 0 { // diff小于零
		this.minV = this.minV - diff
	}
}

func (this *MinStack) Top() int {
	if this.nums[len(this.nums)-1] < 0 {
		return this.minV
	} else {
		return this.nums[len(this.nums)-1] + this.minV
	}
}

func (this *MinStack) GetMin() int {
	if len(this.nums) > 0 {
		return this.minV
	}

	return -1
}
