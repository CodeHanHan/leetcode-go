package offer_47_m_

func maxValue(grid [][]int) int {
	var dp [][]int = make([][]int, len(grid))
	for k, _ := range dp {
		dp[k] = make([]int, len(grid[0]))
	}

	dp[0][0] = grid[0][0]
	for k := 1; k < len(dp[0]); k++ {
		dp[0][k] = dp[0][k-1] + grid[0][k] // 初始化第一行
	}
	for k := 1; k < len(dp); k++ {
		dp[k][0] = dp[k-1][0] + grid[k][0] // 初始化第一列
	}

	for i := 1; i < len(grid); i++ { // 行
		for j := 1; j < len(grid[0]); j++ { // 列
			if dp[i-1][j] > dp[i][j-1] { // 上>左
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else { // 左>上
				dp[i][j] = dp[i][j-1] + grid[i][j]
			}
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]

}
