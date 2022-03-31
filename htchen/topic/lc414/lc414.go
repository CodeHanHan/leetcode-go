package lc414

import (
	"math"
	"sort"
)

func thirdMax1(nums []int) int {
	n := len(nums)
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	cnt := 1
	i := 1
	for ; i < n; i++ {
		if nums[i] != nums[i-1] {
			cnt++
			if cnt == 3 {
				return nums[i]
			}
		}
	}
	return nums[0]
}

func thirdMax2(nums []int) int {
	a, b, c := math.MinInt64, math.MinInt64, math.MinInt64
	for i := 0; i < len(nums); i++ {
		if nums[i] > a {
			a, b, c = nums[i], a, b
		} else if nums[i] > b && nums[i] < a {
			b, c = nums[i], b
		} else if nums[i] < b && nums[i] > c {
			c = nums[i]
		}
	}
	if c == math.MinInt64 {
		return a
	}
	return c
}
