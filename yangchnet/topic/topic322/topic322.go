package topic322

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	dp[0] = 0

	for i := 1; i <= amount; i++ {
		var m int = math.MaxInt64
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != -1 {
				m = min(m, dp[i-coin]+1)
			}
		}
		if m == math.MaxInt64 {
			dp[i] = -1
		} else {
			dp[i] = m
		}
	}

	return dp[amount]
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func coinChange1(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i, _ := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin >= 0 {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}
