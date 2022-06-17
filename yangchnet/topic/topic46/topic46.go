package topic46

func permute(nums []int) [][]int {
	if len(nums) <= 1 { // 数组小于等于1个，直接返回原数组
		return [][]int{nums}
	}

	var ret [][]int = make([][]int, 0)
	var used map[int]bool = make(map[int]bool)

	var fn func(path []int, used map[int]bool)
	fn = func(path []int, used map[int]bool) {
		if len(path) == len(nums) {
			dst := make([]int, len(path))
			copy(dst, path)
			ret = append(ret, dst)
			return
		}

		for _, n := range nums {
			if !used[n] {
				used[n] = true
				path = append(path, n)
				fn(path, used)

				delete(used, n)
				path = path[:len(path)-1]
			}
		}
	}

	fn([]int{}, used)
	return ret
}

func permute1(nums []int) [][]int {
	n := len(nums)
	if n <= 1 {
		return [][]int{nums}
	}

	output := make([]int, n)
	copy(output, nums)

	ans := make([][]int, 0)

	var backtrack func(n, idx int, output []int)
	backtrack = func(n, idx int, output []int) {
		if idx == n {
			tmp := make([]int, n)
			copy(tmp, output)
			ans = append(ans, tmp)

			return
		}

		for i := idx; i < n; i++ {
			output[idx], output[i] = output[i], output[idx]
			backtrack(n, idx+1, output)
			output[idx], output[i] = output[i], output[idx]
		}
	}

	backtrack(n, 0, output)

	return ans
}
