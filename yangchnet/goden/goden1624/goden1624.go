package goden1624

import "sort"

func pairSums(nums []int, target int) [][]int {
	sort.Ints(nums)

	p, q := 0, len(nums)-1

	ret := make([][]int, 0)
	for p < q {
		if nums[p]+nums[q] == target {
			ret = append(ret, []int{nums[p], nums[q]})
			p++
			q--
		} else if nums[p]+nums[q] < target {
			p++
		} else if nums[p]+nums[q] > target {
			q--
		}
	}

	return ret
}
