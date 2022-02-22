package topic33

func search(nums []int, target int) int {
	var l, r int = 0, len(nums) - 1
	var mid = 0
	for l <= r {
		mid = l + (r-l)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= nums[l] { // mid 在左半段
			if target >= nums[l] && target < nums[mid] { // target 也在左边
				r = mid - 1
			} else { // target 在右半段
				l = mid + 1
			}
		} else { // mid 在右半段
			if target > nums[mid] && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
