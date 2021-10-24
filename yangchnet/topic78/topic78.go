package topic78

import "sort"

func subsets(nums []int) [][]int {

	sort.Ints(nums)

	var ret [][]int = make([][]int, 0)

	var fn func(path []int, used map[int]bool, l int, m int)
	fn = func(path []int, used map[int]bool, l int, m int) {
		if len(path) == l {
			dst := make([]int, len(path))
			copy(dst, path)
			ret = append(ret, dst)
			return
		}

		for _, n := range nums {
			if !used[n] {
				if n > m { // 控制程序只往后查找, 避免重复
					path = append(path, n)
					fn(path, used, l, n)

					path = path[:len(path)-1]
				}
			}
		}
	}

	for i := 0; i <= len(nums); i++ {
		fn([]int{}, nil, i, nums[0]-1)
	}

	return ret
}
