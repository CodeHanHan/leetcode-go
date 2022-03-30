package goden0814

func countEval(s string, result int) int {
	n := len(s)
	dp := make([][][2]int, n) // dp[i][j][0/1], s[i:j+1]中计算结果为0/1的个数
	for i := 0; i < n; i++ {
		dp[i] = make([][2]int, n)
	}

	for i := 0; i < n; i += 2 {
		dp[i][i][s[i]-'0']++
	}

	for step := 2; step < n; step += 2 {
		for i := 0; i+step < n; i += 2 {
			for j := i + 1; j < i+step; j += 2 {
				dp[i][i+step][cal(0, 0, s[j])] += dp[i][j-1][0] * dp[j+1][i+step][0]
				dp[i][i+step][cal(0, 1, s[j])] += dp[i][j-1][0]*dp[j+1][i+step][1] + dp[i][j-1][1]*dp[j+1][i+step][0]
				dp[i][i+step][cal(1, 1, s[j])] += dp[i][j-1][1] * dp[j+1][i+step][1]
			}
		}
	}
	return dp[0][n-1][result]
}

func cal(a, b int, op byte) int {
	switch op {
	case '&':
		return a & b
	case '|':
		return a | b
	case '^':
		return a ^ b
	}

	return 0
}
