package lc169

import "sort"

func majorityElement(nums []int) int {
	res, count := 0, 0
	for _, val := range nums {
		if count == 0 {
			res = val
		}
		if val == res {
			count++
		} else {
			count--
		}
	}
	return res
}

func majorityElement_1(nums []int) int {
	length := len(nums)
	sort.Ints(nums)
	return nums[length/2]
}
