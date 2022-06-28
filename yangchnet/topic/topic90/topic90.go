package topic90

import "sort"

func subsetsWithDup(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}

	n := len(nums)
	ans := [][]int{{}, append([]int(nil), nums...)}

	sort.Ints(nums)

	path := []int{}
	vis := make([]bool, n)

	var bp func(idx int, l int) // idx标识path中已经存了多少个数字
	bp = func(idx int, l int) {
		if len(path) == l {
			ans = append(ans, append([]int(nil), path...))
			return
		}

		for i := idx; i < n; i++ {
			if vis[i] || i > 0 && !vis[i-1] && nums[i] == nums[i-1] { // 含重复数字的去重
				continue
			}

			path = append(path, nums[i])
			vis[i] = true
			bp(i, l)
			vis[i] = false
			path = path[:len(path)-1]
		}
	}

	for i := 1; i < n; i++ {
		bp(0, i)
	}

	return ans
}
