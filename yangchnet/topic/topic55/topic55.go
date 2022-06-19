package topic55

func canJump(nums []int) bool {
	n := len(nums)
	reach := 0           // 当前视线范围
	nextReach := nums[0] // 下一步视线范围

	for i := 0; i < n; i++ {
		nextReach = max(nextReach, nums[i]+i)
		if nextReach >= n-1 {
			return true
		}

		if i == reach {
			if nextReach == reach { // 当前视线范围和下一步视线范围相同，也就说走不下去了
				return false
			}
			reach = nextReach
		}
	}

	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
