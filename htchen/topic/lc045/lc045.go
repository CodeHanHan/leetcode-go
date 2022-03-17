package lc045

func jump(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 0
	for i := 1; i < n; i++ {
		dp[i] = n + 1
	}

	for i := 0; i < n; i++ {
		for j := 1; j <= nums[i]; j++ {
			if i+j >= n {
				return dp[n-1]
			}
			dp[i+j] = min(dp[i+j], dp[i]+1)
		}
	}
	return dp[n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
