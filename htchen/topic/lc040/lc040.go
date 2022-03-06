package lc040

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	track := make([]int, 0)
	n := len(candidates)
	sum := 0
	sort.Ints(candidates)

	var dfs func(start int)
	dfs = func(start int) {
		if sum == target {
			ans = append(ans, append([]int{}, track...))
			return
		}
		if sum > target {
			return
		}

		for i := start; i < n; i++ {
			//去重（重要）
			if i > start && candidates[i-1] == candidates[i] {
				continue
			}
			sum += candidates[i]
			track = append(track, candidates[i])
			dfs(i + 1)
			track = track[:len(track)-1]
			sum -= candidates[i]
		}
	}
	dfs(0)
	return ans
}
