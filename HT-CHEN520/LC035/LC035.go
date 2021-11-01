package LC035

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	var mid int
	for left <= right {
		mid = (left + right) / 2
		if target == nums[mid] {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
