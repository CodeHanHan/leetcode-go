package topic034

import "sort"

func searchRange(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			for j := i; j < len(nums); j++ {
				if nums[j] != target {
					return []int{i, j - 1}
				}
			}
			return []int{i, len(nums) - 1}
		}
	}

	return []int{-1, -1}
}

func searchRange_1(nums []int, target int) []int {
	lo, hi := 0, len(nums)-1
	var mid int

	var first int = -1
	for lo <= hi {
		mid = lo + (hi-lo)/2
		if nums[mid] < target {
			lo = mid + 1
		} else if nums[mid] > target {
			hi = mid - 1
		} else {
			if mid == 0 || nums[mid-1] != target {
				first = mid
				break
			} else {
				hi = mid - 1
			}
		}
	}

	if first < 0 {
		return []int{-1, -1}
	}

	for i := first; i <= len(nums); i++ {
		if i >= len(nums) || nums[i] != target {
			return []int{first, i - 1}
		}
	}

	return []int{-1, -1}
}

func searchRange_2(nums []int, target int) []int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}
	rightmost := sort.SearchInts(nums, target+1) - 1
	return []int{leftmost, rightmost}
}
