package dynamic

import (
	"github.com/CodeHanHan/leetcode-go/algo"
)

/*
有几种不同币值的硬币 v1，v2，……，vn（单位是元）
如果要支付 w 元，求最少需要多少个硬币。
比如，我们有 3 种不同的硬币，1 元、3 元、5 元，我们要支付 9 元，最少需要 3 个硬币（3 个 3 元的硬币）。
*/

// 状态转移方程： f(w) = 1 + min{f(w-v1), f(w-v2), ..., f(w-vn)}

func CoinChange(w int, coins []int) int {
	dp := make([]int, w+1)
	for i := 1; i <= w; i++ {
		tmp := make([]int, 0)
		for v := 0; v < len(coins); v++ {
			if i-coins[v] >= 0 {
				tmp = append(tmp, dp[i-coins[v]])
				continue
			}
			break
		}
		dp[i] = 1 + algo.MinN(tmp...)
	}

	return dp[len(dp)-1]
}
