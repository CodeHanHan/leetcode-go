package lc091

func numDecodings(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1] + dp[i]
		}
		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+(s[i-1]-'0')) <= 26 {
			dp[i] = dp[i-2] + dp[i]
		}
	}
	return dp[n]
}
