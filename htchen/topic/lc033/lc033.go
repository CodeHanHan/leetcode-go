package lc033

func search(nums []int, target int) int {
	low, high, mid := 0, len(nums)-1, 0
	for low <= high {
		mid = (low + high) / 2
		if target == nums[mid] {
			return mid
		}
		if nums[mid] >= nums[low] {
			if target >= nums[low] && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}
