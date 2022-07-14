package sort

func quickSort(nums []int) {
	_quickSort(nums, 0, len(nums)-1)
}

func _quickSort(nums []int, left, right int) {
	if left < right {
		pivot := partition(nums, left, right)
		_quickSort(nums, left, pivot-1)
		_quickSort(nums, pivot+1, right)
	}
}

func partition(nums []int, left, right int) int {
	pivot := nums[left]
	for left < right {
		for left < right && nums[right] >= pivot {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] < pivot {
			left++
		}
		nums[right] = nums[left]
	}

	nums[left] = pivot

	return left
}
