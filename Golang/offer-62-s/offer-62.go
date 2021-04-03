package offer_62_s

// 暴力法，此解法超时
func lastRemaining(n int, m int) int {
	nums := make([]int, n)
	for i, _ := range nums {
		nums[i] = i
	}

	i, count := 0, 1

	for len(nums) > 1 {
		if i >= len(nums) {
			i = 0
		}
		if count == m {
			nums = append(nums[:i], nums[i+1:]...)
			count = 1
			continue
		}
		i++
		count++
	}
	return nums[0]
}

// 数学解法
func lastRemaining1(n int, m int) int {
	ret := 0
	// 由分析可知，最后一轮剩余两个数字
	for i := 2; i <= n; i++ {
		ret = (ret + m) % i
	}
	return ret
}
