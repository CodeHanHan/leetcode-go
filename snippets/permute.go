package snippets

// 全排列

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

// 有重复数字的全排列
func permuteUnique1(nums []int) (ans [][]int) {
	// TODO implement
	return nil
}
