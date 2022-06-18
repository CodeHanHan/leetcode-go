package topic45

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	reach := 0           // 当前视线范围
	nextreach := nums[0] // 下一步视线范围
	step := 0

	for i := 0; i < len(nums); i++ {
		nextreach = max(nextreach, nums[i]+i)
		if nextreach >= len(nums)-1 {
			return step + 1
		}

		if i == reach { // 当走到当前视线尽头，需要再走一步
			step++
			reach = nextreach // 更新视线范围
		}
	}

	return step
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
