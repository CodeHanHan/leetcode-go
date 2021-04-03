package offer_59_II_m

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */

type MaxQueue struct {
	nums []int // 存储所有进队列的元素，负责pop
	max  []int // 非严格单调减存储元素，负责返回max
}

func Constructor() MaxQueue {
	return MaxQueue{
		nums: make([]int, 0),
		max:  make([]int, 0),
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.max) == 0 {
		return -1
	}
	return this.max[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.nums = append(this.nums, value)
	for len(this.max) > 0 && this.max[len(this.max)-1] < value {
		this.max = this.max[:len(this.max)-1]
	}
	this.max = append(this.max, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.nums) == 0 {
		return -1
	}
	pop := this.nums[0]
	this.nums = this.nums[1:]
	if this.max[0] == pop {
		this.max = this.max[1:]
	}
	return pop
}

/*
["MaxQueue","pop_front","pop_front","pop_front","pop_front","pop_front","push_back","max_value","push_back","max_value"]
[[],[],[],[],[],[],[15],[],[9],[]]
*/
