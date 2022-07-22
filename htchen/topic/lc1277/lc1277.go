package lc1277

func countSquares(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	cnt := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = matrix[i][j]
			} else if matrix[i][j] == 0 {
				dp[i][j] = 0
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
			cnt += dp[i][j]
		}
	}
	return cnt
}

func min(x ...int) int {
	min := x[0]
	for _, v := range x {
		if v < min {
			min = v
		}
	}
	return min
}
