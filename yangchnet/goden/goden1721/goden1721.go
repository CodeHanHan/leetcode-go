package goden1721

func trap(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}
	rms := make([]int, n)
	lms := make([]int, n)

	rm := height[n-1]
	for i := n - 1; i >= 0; i-- {
		if height[i] > rm {
			rm = height[i]
		}
		rms[i] = rm
	}

	lm := height[0]
	for i := 0; i < n; i++ {
		if height[i] > lm {
			lm = height[i]
		}
		lms[i] = lm
	}

	count := 0
	for i := 1; i < n-1; i++ {
		count += min(lms[i], rms[i]) - height[i]
	}

	return count
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
