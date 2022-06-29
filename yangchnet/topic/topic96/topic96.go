package topic96

func numTrees(n int) int {
	if n <= 2 {
		return n
	}

	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 1, 1, 2

	for i := 3; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}

	return dp[n]
}
