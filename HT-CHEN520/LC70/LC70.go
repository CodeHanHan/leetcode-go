package LC70

func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}

	var dp []int = make([]int, n+1)

	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func climbStairs_1(n int) int {
	if n <= 1 {
		return 1
	}

	var p, q, r int = 1, 1, -1

	for i := 2; i <= n; i++ {
		r = p + q
		p = q
		q = r
	}

	return r
}
