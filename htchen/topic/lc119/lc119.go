package lc119

func getRow(rowIndex int) []int {
	dp := make([][]int, rowIndex+1)
	for i := 0; i < rowIndex+1; i++ {
		dp[i] = make([]int, i+1)
		dp[i][0], dp[i][i] = 1, 1
		for j := 1; j < i; j++ {
			dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
		}
	}
	return dp[rowIndex]
}
