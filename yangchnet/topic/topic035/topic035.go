package topic035

func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if target <= nums[i] {
			return i
		}
	}
	return len(nums)
}

// 找到升序无重复数组中等于目标值的数的位置或大于目标值的第一个数的位置
func searchInsert_1(nums []int, target int) int {
	lo, hi := 0, len(nums)-1

	var mid int
	for lo <= hi {
		mid = lo + (hi-lo)/2
		if nums[mid] < target {
			lo = mid + 1
		} else if nums[mid] > target {
			hi = mid - 1
			if mid == 0 || mid-1 >= 0 && nums[mid-1] < target {
				return mid
			}
		} else {
			return mid
		}
	}

	return len(nums)
}
