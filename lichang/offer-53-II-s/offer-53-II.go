package offer_53_II_s_

func missingNumber(nums []int) int {
	left, right := 0, len(nums)-1
	var mid int
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
