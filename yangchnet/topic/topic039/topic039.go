package topic039

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	ans := make([][]int, 0)
	existed := make(map[string]bool)

	var fn func(tmp []int, sum int)
	fn = func(tmp []int, sum int) {
		if sum > target {
			return
		}

		if sum == target {
			cp := make([]int, len(tmp))
			copy(cp, tmp)
			sort.Ints(cp)
			if !existed[fmt.Sprintf("%v", cp)] {
				ans = append(ans, cp)
				existed[fmt.Sprintf("%v", cp)] = true
			}
			return
		}

		for i := 0; i < len(candidates); i++ {
			if target-sum >= candidates[i] {
				fn(append(tmp, candidates[i]), sum+candidates[i])
			} else {
				break
			}
		}
	}

	fn([]int{}, 0)

	return ans
}

func combinationSum_1(candidates []int, target int) (ans [][]int) {
	comb := []int{}

	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}

		// 直接跳过
		dfs(target, idx+1)

		// 选择当前数
		if target-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx) // 每个数字可以被无限制重复选取，因此搜索的下标仍为 idx
			comb = comb[:len(comb)-1]
		}
	}

	dfs(target, 0)
	return
}
