package goden1617

// 动态规划
func maxSubArray(nums []int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}

	var a int
	a = nums[0]

	m := a
	for i := 1; i < n; i++ {
		a = max(a+nums[i], nums[i])
		if m < a {
			m = a
		}
	}

	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
