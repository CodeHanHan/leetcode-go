package topic91

import "fmt"

func numDecodings(s string) int {
	n := len(s)
	if n <= 0 || (s[0]-'0') == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	ss := fmt.Sprintf(" %s", s)

	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= n; i++ {
		if ss[i]-'0' == 0 && ss[i-1]-'0' == 0 {
			return 0
		}

		if ss[i]-'0' > 0 {
			dp[i] += dp[i-1]
		}
		if ss[i-1]-'0' > 0 && 10*(ss[i-1]-'0')+(ss[i]-'0') <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}
