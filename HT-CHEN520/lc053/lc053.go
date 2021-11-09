package lc053

func maxSubArray(nums []int) int {
	len := len(nums)
	f := nums[0]
	maxSum := nums[0]
	for i := 1; i < len; i++ {
		f = max(f+nums[i], nums[i])
		if f > maxSum {
			maxSum = f
		}
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
