package lc034

func searchRange(nums []int, target int) []int {
	len := len(nums)
	low, high, mid := 0, len-1, 0
	targetMid := -1
	for low <= high {
		mid = (low + high) / 2
		if nums[mid] == target {
			targetMid = mid
			break
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	leftIndex, rightIndex := 0, 0
	if targetMid == -1 {
		return []int{-1, -1}
	} else {
		for i := targetMid; i >= 0; i-- {
			if nums[i] == target {
				leftIndex = i
			} else {
				break
			}
		}
		for j := targetMid; j < len; j++ {
			if nums[j] == target {
				rightIndex = j
			} else {
				break
			}
		}

	}
	return []int{leftIndex, rightIndex}
}
