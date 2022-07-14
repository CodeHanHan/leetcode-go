package lc042

func trap(height []int) int {
	n := len(height)
	res := 0
	left := make([]int, n)
	left[0] = height[0]
	for i := 1; i < n; i++ {
		left[i] = max(left[i-1], height[i])
	}

	right := make([]int, n)
	right[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i])
	}
	for i, v := range height {
		res += min(left[i], right[i]) - v
	}
	return res
}

func trap1(height []int) int {
	n := len(height)
	left, right := 0, n-1
	leftMax, rightMax := 0, 0
	res := 0
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if leftMax < rightMax {
			res += leftMax - height[left]
			left++
		} else {
			res += rightMax - height[right]
			right--
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
