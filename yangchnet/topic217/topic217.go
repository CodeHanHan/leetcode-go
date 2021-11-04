package topic217

func containsDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	existed := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		if _, ok := existed[nums[i]]; ok {
			return true
		}
		existed[nums[i]] = true
	}

	return false
}
