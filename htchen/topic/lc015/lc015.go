package lc015

import "sort"

func threeSum(nums []int) [][]int {
	len := len(nums)
	ans := [][]int{}
	if nums == nil || len < 3 {
		return ans
	}
	sort.Ints(nums)
	for i := 0; i < len; i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res := []int{nums[i], nums[left], nums[right]}
				ans = append(ans, res)
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return ans
}
