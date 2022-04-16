package lc514

import (
	"math"
)

func findRotateSteps(ring string, key string) int {
	indexMap := make([][]int, 26)
	for i, c := range ring {
		indexMap[c-'a'] = append(indexMap[c-'a'], i)
	}

	dp := make([][]int, len(ring))
	for i, _ := range dp {
		dp[i] = make([]int, len(key))
		for j, _ := range dp[i] {
			dp[i][j] = -1
		}
	}

	var dfs func(ringI, keyI int) int
	dfs = func(ringI, keyI int) int {
		if keyI == len(key) {
			return 0
		}
		if dp[ringI][keyI] != -1 {
			return dp[ringI][keyI]
		}

		cur := key[keyI]
		curInt := cur - 'a'
		res := math.MaxInt32
		for _, targetI := range indexMap[curInt] {
			d1 := abs(ringI - targetI)
			d2 := len(ring) - d1
			curMin := min(d1, d2)
			res = min(res, curMin+dfs(targetI, keyI+1))
		}
		dp[ringI][keyI] = res
		return res
	}
	return len(key) + dfs(0, 0)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
