package lc057

func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	n := len(intervals)

	left, right := newInterval[0], newInterval[1]
	i := 0
	for i < n && intervals[i][1] < left {
		res = append(res, intervals[i])
		i++
	}
	for i < n && intervals[i][0] <= right {
		left = min(intervals[i][0], left)
		right = max(intervals[i][1], right)
		i++
	}

	res = append(res, []int{left, right})
	for i < n {
		res = append(res, intervals[i])
		i++
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
