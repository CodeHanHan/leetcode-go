package lc354

import "sort"

func maxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)
	if n == 0 {
		return 0
	}

	sort.Slice(envelopes, func(i, j int) bool {
		first, second := envelopes[i], envelopes[j]
		return first[0] < second[0] || first[0] == second[0] && first[1] > second[1]
	})

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if envelopes[i][1] > envelopes[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	return max(dp...)
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a {
		if res < v {
			res = v
		}
	}
	return res
}
