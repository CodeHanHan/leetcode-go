package lc078

func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	track := make([]int, 0)

	var dfs func(start int)
	dfs = func(start int) {
		ans = append(ans, append([]int{}, track...))
		for i := start; i < len(nums); i++ {
			track = append(track, nums[i])
			dfs(i + 1)
			track = track[:len(track)-1]
		}
	}
	dfs(0)
	return ans
}
