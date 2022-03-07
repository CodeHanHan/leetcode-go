package lc077

func combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	track := make([]int, 0)

	var dfs func(start int)
	dfs = func(start int) {
		if len(track) == k {
			ans = append(ans, append([]int{}, track...))
			return
		}
		for i := start; i <= n; i++ {
			track = append(track, i)
			dfs(i + 1)
			track = track[:len(track)-1]
		}
	}
	dfs(1)
	return ans
}
