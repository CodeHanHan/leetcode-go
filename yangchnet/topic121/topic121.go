package topic121

import "math"

// 一次遍历
func maxProfit_1(prices []int) int {
	var minPrice int = math.MaxInt32
	var profit int = 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > profit {
			profit = prices[i] - minPrice
		}
	}

	return profit
}

// 动态规划解法， 时间超时
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	dp := make([]int, len(prices))
	dp[0], dp[1] = 0, prices[1]-prices[0]

	for i := 2; i < len(prices); i++ {
		max := math.MinInt32
		for j := 0; j < i; j++ {
			if prices[i]-prices[j] > max {
				max = prices[i] - prices[j]
			}
		}

		if dp[i-1] < max {
			dp[i] = max
		} else {
			dp[i] = dp[i-1]
		}
	}

	if dp[len(dp)-1] < 0 {
		return 0
	}
	return dp[len(dp)-1]
}
