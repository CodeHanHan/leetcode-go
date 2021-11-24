package lc120

import (
	"math"
)

func minimumTotal(triangle [][]int) int {
	len := len(triangle)
	dp := make([][]int, len)
	for i := 0; i < len; i++ {
		dp[i] = make([]int, len)
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < len; i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
		}
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}
	res := math.MaxInt32
	for i := 0; i < len; i++ {
		res = min(res, dp[len-1][i])
	}
	return res
}

func min(x, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}
