package topic41

// 不限制空间
func firstMissingPositive(nums []int) int {
	if len(nums) <= 0 {
		return 1
	}
	existed := make(map[int]bool)
	var max int
	max = nums[0]

	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 {
			continue
		}
		existed[nums[i]] = true

		if max < nums[i] {
			max = nums[i]
		}
	}

	for i := 1; i <= max; i++ {
		if _, ok := existed[i]; !ok {
			return i
		}
	}

	return max + 1
}

func firstMissingPositive1(nums []int) int {
	n := len(nums)
	for i := 0; i < len(nums); i++ {
		// for nums[i] > 0 && nums[i] < n && nums[nums[i]-1] != nums[i] {
		for nums[i] > 0 && nums[i] < n && nums[nums[i]-1] != nums[i] { // nums[nums[i]-1] == nums[i]代表一个数字放在了正确的位置上
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return len(nums) + 1
}
