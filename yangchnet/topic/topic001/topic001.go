package topic001

// 双重循环，暴力求解
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func twoSum_1(nums []int, target int) []int {
	wanted := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if pre, ok := wanted[nums[i]]; ok {
			return []int{pre, i}
		}

		wanted[target-nums[i]] = i
	}

	return []int{}
}
