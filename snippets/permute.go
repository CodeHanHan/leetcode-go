package snippets

import "sort"

// 全排列

// topic46
// 无重复数字的全排列
func permute(nums []int) [][]int {
	n := len(nums)

	if n <= 1 {
		return [][]int{nums}
	}

	ans := make([][]int, 0)

	output := append([]int{}, append([]int(nil), nums...)...)

	var fn func(idx int, output []int) // idx 标识填充到第几个数字
	fn = func(idx int, output []int) {
		if idx == n {
			ans = append(ans, append([]int(nil), output...))
			return
		}

		// 有n个空位，往空位中放入数字，每次迭代是固定了idx之前的所有数字，然后将后面所有未用过的数字往这个位置放一次
		for i := idx; i < n; i++ {
			output[idx], output[i] = output[i], output[idx]
			fn(idx+1, output)
			output[idx], output[i] = output[i], output[idx]
		}
	}

	fn(0, output)

	return ans

}

// topic47
// 有重复数字的全排列
func permuteUnique(nums []int) (ans [][]int) {
	n := len(nums)
	if n <= 1 {
		return [][]int{nums}
	}

	sort.Ints(nums)
	perm := make([]int, 0)

	vis := make([]bool, n)

	var fn func(idx int)
	fn = func(idx int) {
		if idx == n {
			ans = append(ans, append([]int(nil), perm...))
			return
		}

		for i, v := range nums {
			// 去重
			if vis[i] || i > 0 && !vis[i-1] && v == nums[i-1] {
				continue
			}

			perm = append(perm, v)
			vis[i] = true
			fn(idx + 1)
			vis[i] = false
			perm = perm[:len(perm)-1]
		}
	}

	fn(0)

	return ans
}

func permuteUnique1(nums []int) [][]int {
	ans := make([][]int, 0)
	if len(nums) <= 0 {
		return [][]int{}
	}

	var fn func(idx int)
	fn = func(idx int) {
		if idx >= len(nums) {
			ans = append(ans, append([]int(nil), nums...))
			return
		}

		used := make(map[int]bool)
		for i := idx; i < len(nums); i++ {

			if _, ok := used[nums[i]]; ok {
				continue
			}

			used[nums[i]] = true
			nums[idx], nums[i] = nums[i], nums[idx]
			fn(idx + 1)
			nums[idx], nums[i] = nums[i], nums[idx]
		}
	}

	fn(0)

	return ans

}
