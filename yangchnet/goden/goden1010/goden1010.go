package goden1010

type StreamRank struct {
	nums []int
}

func Constructor() StreamRank {
	return StreamRank{
		nums: make([]int, 50001),
	}
}

func (this *StreamRank) Track(x int) {
	// 1. 为新插入的x设置秩; 2. 更新大于x的秩
	this.nums[x]++
}

func (this *StreamRank) GetRankOfNumber(x int) int {
	sum := 0
	for i := 0; i <= x; i++ {
		sum += this.nums[i]
	}

	return sum
}

/**
 * Your StreamRank object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Track(x);
 * param_2 := obj.GetRankOfNumber(x);
 */
