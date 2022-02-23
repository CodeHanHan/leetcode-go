package lc217

import "sort"

func containsDuplicate(nums []int) bool {
	m := make(map[int]int, len(nums))

	for i, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = i
	}
	return false
}

func containsDuplicate_1(nums []int) bool {
	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			return true
		}
	}
	return false
}
