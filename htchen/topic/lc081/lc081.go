package lc081

func search(nums []int, target int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	if n == 1 {
		return target == nums[0]
	}

	left, right, mid := 0, n-1, 0
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			return true
		}

		if nums[mid] == nums[left] && nums[mid] == nums[right] {
			left++
			right--
		} else if nums[mid] >= nums[left] {
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return false
}
