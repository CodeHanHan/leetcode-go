package lc016

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	len := len(nums)
	closest := math.MaxInt32

	update := func(sum int) {
		if math.Abs(float64(sum-target)) < math.Abs(float64(closest-target)) {
			closest = sum
		}
	}

	for i := 0; i < len; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j, k := i+1, len-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return target
			}

			update(sum)
			if sum < target {
				j++
			} else {
				k--
			}
		}
	}
	return closest
}
