package lc135

func candy(ratings []int) int {
	n := len(ratings)
	left := make([]int, n)
	right := make([]int, n)
	for i := 0; i < n; i++ {
		left[i] = 1
		right[i] = 1
	}

	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		}
	}
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			right[i] = right[i+1] + 1
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans += max(left[i], right[i])
	}
	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
