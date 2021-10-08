package topic16

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	if len(nums) == 3 {
		return nums[0] + nums[1] + nums[2]
	}

	sort.Ints(nums)

	closest := math.MaxInt32
	for i := 0; i < len(nums)-2; i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			sum := sum3(nums, i, l, r)
			if abs(sum, target) <= abs(closest, target) {
				closest = sum
			}
			if sum-target == 0 {
				return closest
			} else if sum-target > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return closest
}

func abs(x, y int) int {
	if x-y <= 0 {
		return y - x
	}
	return x - y
}

func sum3(nums []int, i, j, k int) int {
	return nums[i] + nums[j] + nums[k]
}
