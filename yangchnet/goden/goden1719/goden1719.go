package goden1719

func missingTwo(nums []int) []int {
	ret := make([]int, 0)
	nums = append(nums, -1)
	nums = append(nums, -1)

	for i := 0; i < len(nums); i++ {
		for nums[i] != -1 && nums[i] != i+1 {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == -1 {
			ret = append(ret, i+1)
		}
	}

	return ret
}
