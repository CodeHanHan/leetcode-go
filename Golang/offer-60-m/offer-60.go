package offer_60_m

import "math"

func dicesProbability(n int) []float64 {
	dp := make([][]int, n+1)
	for i, _ := range dp {
		dp[i] = make([]int, 6*n+1)
	}

	for i := 1; i <= 6; i++ {
		dp[1][i] = 1
	}

	for i := 2; i <= n; i++ {
		for j := i; j <= i*6; j++ { // 翻译状态转移方程
			for cur := 1; cur <= 6; cur++ {
				if j-cur <= 0 {
					break
				}
				dp[i][j] += dp[i-1][j-cur]
			}
		}
	}

	all := math.Pow(float64(6), float64(n))
	res := make([]float64, 0)
	for i := n; i <= 6*n; i++ {
		res = append(res, float64(dp[n][i])/all)
	}

	return res
}

// 空间优化
func dicesProbability1(n int) []float64 {
	dp := make([]int, 6*n+1)

	for i := 1; i <= 6; i++ {
		dp[i] = 1
	}

	for i := 2; i <= n; i++ {
		for j := 6 * i; j >= i; j-- { // 从后往前，不可从前往后
			dp[j] = 0
			for cur := 1; cur <= 6; cur++ {
				if j-cur < i-1 {
					break
				}
				dp[j] += dp[j-cur]
			}
		}
	}

	all := math.Pow(float64(6), float64(n))
	res := make([]float64, 0)
	for i := n; i <= 6*n; i++ {
		res = append(res, float64(dp[i])/all)
	}

	return res
}
