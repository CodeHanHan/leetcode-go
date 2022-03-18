package goden0804

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}
	res := make([][]int, 0)

	path := make([]int, 0)
	var fn func(idx int)
	fn = func(idx int) {
		if idx >= len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		path = append(path, nums[idx])
		fn(idx + 1)

		path = path[:len(path)-1]
		fn(idx + 1)
	}

	fn(0)

	return res
}
