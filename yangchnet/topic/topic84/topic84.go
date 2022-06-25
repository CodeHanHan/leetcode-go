package topic84

func largestRectangleArea(heights []int) int {
	n := len(heights)
	leftS := make([]int, n)
	rightS := make([]int, n)

	s := make([][]int, 0)
	for i := 0; i < n; i++ {
		for len(s) > 0 && s[len(s)-1][0] >= heights[i] {
			s = s[:len(s)-1]
		}
		s = append(s, []int{heights[i], i})

		if len(s) > 1 {
			leftS[i] = s[len(s)-2][1]
		} else {
			leftS[i] = -1
		}
	}

	s = [][]int{}
	for i := n - 1; i >= 0; i-- {
		for len(s) > 0 && s[len(s)-1][0] >= heights[i] {
			s = s[:len(s)-1]
		}
		s = append(s, []int{heights[i], i})

		if len(s) > 1 {
			rightS[i] = s[len(s)-2][1]
		} else {
			rightS[i] = n
		}

	}

	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, (rightS[i]-leftS[i]-1)*heights[i])
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
