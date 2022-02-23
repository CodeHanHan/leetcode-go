package lc080

func removeDuplicates(nums []int) int {
	len := len(nums)
	if len <= 2 {
		return len
	}

	slow, fast := 2, 2
	for fast < len {
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
