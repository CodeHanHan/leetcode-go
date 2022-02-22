package topic169

func majorityElement(nums []int) int {
	var tickets, candidate int = 1, nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] == candidate {
			tickets += 1
		} else {
			if tickets > 0 {
				tickets -= 1
			} else {
				candidate = nums[i]
				tickets = 1
			}
		}
	}

	return candidate
}
