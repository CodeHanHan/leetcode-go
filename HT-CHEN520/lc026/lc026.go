package lc26

func removeDuplicates(nums []int) int {
	len := len(nums)

	if len == 0 {
		return 0
	}

	i, j := 0, 0
	for j < len {
		if nums[i] == nums[j] {
			j++
		} else {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
