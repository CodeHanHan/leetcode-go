package topic154

func findMin(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mid := l + (r-l)/2

		if nums[mid] > nums[r] { // min在mid右边
			l = mid + 1
		} else if nums[mid] < nums[r] { // min在mid左边
			r = mid
		} else if nums[mid] == nums[r] { // 不能确定
			r--
		}
	}

	return min(nums[l], nums[r])
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
