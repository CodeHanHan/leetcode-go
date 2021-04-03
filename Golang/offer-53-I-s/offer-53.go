package offer_53_s_

// O(n)
func search(nums []int, target int) int {
	count := 0
	for _, v := range nums {
		if v > target {
			break
		}
		if v == target {
			count++
		}
	}
	return count
}

// O(logn)
func search1(nums []int, target int) int {
	count := 0
	if len(nums) < 1 {
		return count
	}
	left, right := 0, len(nums)-1
	mid := left + (right-left)/2
	for left <= right && nums[mid] != target {
		if nums[mid] < target {
			left = mid + 1
			mid = left + (right-left)/2
		} else {
			right = mid - 1
			mid = left + (right-left)/2
		}
	}
	if left <= right && nums[mid] == target {
		count++
		for i := mid + 1; i < len(nums); i++ {
			if nums[i] > target {
				break
			}
			if nums[i] == target {
				count++
			}
		}
		for i := mid - 1; i >= 0; i-- {
			if nums[i] < target {
				break
			}
			if nums[i] == target {
				count++
			}
		}
	}
	return count
}
