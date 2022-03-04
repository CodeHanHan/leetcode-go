package goden0301

type TripleInOne struct {
	nums      []int
	stackSize int
	ptrs      map[int]int
	// 第i段：nums[i*stackSize : (i+1)*stackSize], i从0开始
	// 第i段空：ptr==i*stackSize
	// 第i段满：ptr==(i+1)*stackSize
	// 第i段Peek:  nums[ptrs[i]-1]
}

func Constructor(stackSize int) TripleInOne {
	t := TripleInOne{
		nums:      make([]int, stackSize*3),
		stackSize: stackSize,
		ptrs:      make(map[int]int),
	}

	for i := 0; i < 3; i++ {
		t.ptrs[i] = stackSize * i
	}

	return t
}

func (this *TripleInOne) Push(stackNum int, value int) {
	if this.ptrs[stackNum] >= (stackNum+1)*this.stackSize {
		return
	}
	this.nums[this.ptrs[stackNum]] = value
	this.ptrs[stackNum]++
}

func (this *TripleInOne) Pop(stackNum int) int {
	if this.IsEmpty(stackNum) {
		return -1
	}

	this.ptrs[stackNum]--
	return this.nums[this.ptrs[stackNum]]
}

func (this *TripleInOne) Peek(stackNum int) int {
	if this.IsEmpty(stackNum) {
		return -1
	}

	return this.nums[this.ptrs[stackNum]-1]
}

func (this *TripleInOne) IsEmpty(stackNum int) bool {
	return this.ptrs[stackNum] == stackNum*this.stackSize
}
