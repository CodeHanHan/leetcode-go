package topic47

import "sort"

func permuteUnique(nums []int) [][]int {
	n := len(nums)
	if n <= 1 {
		return [][]int{nums}
	}

	output := make([]int, n)
	copy(output, nums)
	sort.Ints(output)

	ans := make([][]int, 0)

	var backtrack func(n, idx int, output []int)
	backtrack = func(n, idx int, output []int) {
		if idx == n {
			tmp := make([]int, n)
			copy(tmp, output)
			ans = append(ans, tmp)

			return
		}

		used := make(map[int]bool)
		for i := idx; i < n; i++ {
			if _, ok := used[output[i]]; ok {
				continue
			}
			output[idx], output[i] = output[i], output[idx]
			backtrack(n, idx+1, output)
			output[idx], output[i] = output[i], output[idx]
			used[output[i]] = true
		}
	}

	backtrack(n, 0, output)

	return ans
}

func permuteUnique1(nums []int) (ans [][]int) {
	sort.Ints(nums)

	n := len(nums)
	perm := []int{}

	vis := make([]bool, n)

	var backtrack func(int)
	backtrack = func(idx int) {
		if idx == n {
			ans = append(ans, append([]int(nil), perm...))
			return
		}

		for i, v := range nums {
			if vis[i] || i > 0 && !vis[i-1] && v == nums[i-1] {
				continue
			}

			perm = append(perm, v)
			vis[i] = true
			backtrack(idx + 1)
			vis[i] = false
			perm = perm[:len(perm)-1]
		}
	}

	backtrack(0)

	return
}
