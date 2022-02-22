package topic031

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	i, j, k := len(nums)-2, len(nums)-1, len(nums)-1

	for i >= 0 && nums[i] >= nums[j] { //从后向前找到第一个相邻升序对
		i--
		j--
	}

	if i >= 0 { // 找到后半部分中比nums[i]大的最小值
		for nums[i] >= nums[k] {
			k--
		}
		nums[i], nums[k] = nums[k], nums[i]
	}

	for i, j := j, len(nums)-1; i < j; i, j = i+1, j-1 { // 反转使其变成升序
		nums[i], nums[j] = nums[j], nums[i]
	}

}
