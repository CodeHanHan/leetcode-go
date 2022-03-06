package lc047

import "sort"

func permuteUnique(nums []int) [][]int {
	ans := make([][]int, 0)
	track := make([]int, 0)
	n := len(nums)
	used := map[int]bool{}
	sort.Ints(nums)

	var dfs func()
	dfs = func() {
		if len(track) == n {
			ans = append(ans, append([]int{}, track...))
			return
		}

		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			if i > 0 && used[i-1] == false && nums[i-1] == nums[i] {
				continue
			}
			used[i] = true
			track = append(track, nums[i])
			dfs()
			used[i] = false
			track = track[:len(track)-1]
		}
	}
	dfs()
	return ans
}
