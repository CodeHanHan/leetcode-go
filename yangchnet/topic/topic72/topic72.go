package topic72

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m*n == 0 {
		return m + n
	}

	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m+1; i++ {
		dp[i][0] = i
	}

	for j := 0; j < n+1; j++ {
		dp[0][j] = j
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = 1 + minN(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]-1)
			} else {
				dp[i][j] = 1 + minN(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
			}
		}
	}

	return dp[m][n]
}

func minN(nums ...int) int {
	min := nums[0]
	for _, v := range nums {
		if min > v {
			min = v
		}
	}

	return min
}
