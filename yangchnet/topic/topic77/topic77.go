package topic77

func combine(n int, k int) [][]int {
	ans := make([][]int, 0)

	path := []int{}

	var bp func(path []int, used map[int]bool)
	bp = func(path []int, used map[int]bool) {
		if len(path) == k {
			ans = append(ans, append([]int(nil), path...))
			return
		}

		var i int
		if len(path) <= 0 {
			i = 1
		} else {
			i = path[len(path)-1]
		}

		for ; i <= n; i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, i)
			bp(path, used)
			used[i] = false
			path = path[:len(path)-1]
		}
	}

	used := make(map[int]bool)
	bp(path, used)

	return ans
}
