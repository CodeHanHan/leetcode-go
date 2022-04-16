package goden1716

// func massage(nums []int) int {
// 	n := len(nums)
// 	if n <= 0 {
// 		return 0
// 	}
// 	dp := make([][]int, 2)
// 	for i := range dp {
// 		dp[i] = make([]int, n)
// 	}

// 	dp[0][0] = 0
// 	dp[1][0] = nums[0]
// 	for i := 1; i < n; i++ {
// 		dp[0][i] = max(dp[0][i-1], dp[1][i-1])
// 		dp[1][i] = dp[0][i-1] + nums[i]
// 	}

// 	return max(dp[0][n-1], dp[1][n-1])
// }

func massage(nums []int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}

	dp0 := 0
	dp1 := nums[0]
	for i := 1; i < n; i++ {
		dp0, dp1 = max(dp0, dp1), dp0+nums[i]
	}

	return max(dp0, dp1)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
