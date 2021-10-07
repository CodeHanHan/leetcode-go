package topic15

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	sort.Ints(nums)

	ret := make([][]int, 0)

	for i := 0; i < len(nums) && nums[i] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			if nums[i]+nums[l]+nums[r] == 0 {
				ret = append(ret, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l += 1
				}
				for l < r && nums[r] == nums[r-1] {
					r -= 1
				}
				l++
				r--
			} else if nums[i]+nums[l]+nums[r] < 0 {
				l++
			} else if nums[i]+nums[l]+nums[r] > 0 {
				r--
			}
		}
	}
	return ret
}
