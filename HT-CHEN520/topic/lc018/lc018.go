package lc018

import "sort"

func fourSum(nums []int, target int) [][]int {
	len := len(nums)
	ans := [][]int{}
	if nums == nil || len < 4 {
		return ans
	}
	sort.Ints(nums)
	for i := 0; i < len-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			left, right := j+1, len-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					res := []int{nums[i], nums[j], nums[left], nums[right]}
					ans = append(ans, res)
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if sum > target {
					right--
				} else {
					left++
				}
			}
		}
	}
	return ans
}
