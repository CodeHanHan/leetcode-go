package lc581

import (
	"math"
	"sort"
)

func findUnsortedSubarray(nums []int) int {
	if sort.IntsAreSorted(nums) {
		return 0
	}
	sorted := append([]int(nil), nums...)
	sort.Ints(sorted)
	left, right := 0, len(nums)-1
	for nums[left] == sorted[left] {
		left++
	}
	for nums[right] == sorted[right] {
		right--
	}
	return right - left + 1
}

func findUnsortedSubarray1(nums []int) int {
	n := len(nums)
	min, max := math.MaxInt64, math.MinInt64
	left, right := -1, -1
	for i, num := range nums {
		if max > num {
			right = i
		} else {
			max = num
		}
		if min < nums[n-i-1] {
			left = n - i - 1
		} else {
			min = nums[n-i-1]
		}
	}
	if right == -1 {
		return 0
	}
	return right - left + 1
}
