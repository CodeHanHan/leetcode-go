package topic912

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)

	return nums
}

func quickSort(nums []int, left, right int) {
	if left < right {
		pivotIdx := partition(nums, left, right)
		quickSort(nums, left, pivotIdx-1)
		quickSort(nums, pivotIdx+1, right)
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
