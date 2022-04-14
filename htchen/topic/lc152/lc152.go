package lc152

import "math"

func maxProduct(nums []int) int {
	n := len(nums)
	imax, imin, res := 1, 1, math.MinInt64
	for i := 0; i < n; i++ {
		if nums[i] < 0 {
			imax, imin = imin, imax
		}
		imax = max(imax*nums[i], nums[i])
		imin = min(imin*nums[i], nums[i])

		res = max(imax, res)
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
