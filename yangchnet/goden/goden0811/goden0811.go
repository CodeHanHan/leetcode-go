package goden0811

func waysToChange(n int) int {
	coins := []int{1, 5, 10, 25}
	dp := make([]int, n+1)
	dp[0] = 1

	for _, coin := range coins {
		for i := coin; i <= n; i++ {
			dp[i] = (dp[i] + dp[i-coin]) % 1000000007
		}
	}

	return dp[n]
}
