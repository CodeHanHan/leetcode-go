package topic042

func trap(height []int) int {
	left := make([]int, len(height))
	max := height[0]
	for i := 0; i < len(height); i++ {
		if height[i] > max {
			max = height[i]
		}
		left[i] = max
	}

	right := make([]int, len(height))
	max = height[len(height)-1]
	for i := len(height) - 1; i >= 0; i-- {
		if height[i] > max {
			max = height[i]
		}
		right[i] = max
	}

	capacity := 0
	for i := 0; i < len(height); i++ {
		capacity += min(left[i], right[i]) - height[i]
	}

	return capacity
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
