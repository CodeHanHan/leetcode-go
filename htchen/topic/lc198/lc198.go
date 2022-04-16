package lc198

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	first, second, temp := nums[0], max(nums[0], nums[1]), 0
	for i := 2; i < n; i++ {
		temp = first
		first = second
		second = max(temp+nums[i], second)
	}
	return second
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
