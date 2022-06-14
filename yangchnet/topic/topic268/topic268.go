package topic268

func missingNumber(nums []int) int {
	var m = nums[0]
	for i := 1; i < len(nums); i++ {
		m ^= nums[i]
	}

	for i := 0; i <= len(nums); i++ {
		m ^= i
	}

	return m
}
