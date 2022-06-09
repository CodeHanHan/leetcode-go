package topic81

func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return true
		}

		if nums[mid] > nums[l] { // mid落在左半段
			if target >= nums[l] && target < nums[mid] { // target也落在左半段
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else if nums[mid] < nums[l] { // mid 落在右半段
			if target <= nums[r] && target > nums[mid] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else if nums[mid] == nums[l] {
			l++
		}
	}

	return false
}
