package lc300

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	for i, _ := range dp {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	return max(dp...)
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a {
		if res < v {
			res = v
		}
	}
	return res
}
