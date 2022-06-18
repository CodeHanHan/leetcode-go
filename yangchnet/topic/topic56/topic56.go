package topic56

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] >= intervals[j][1] {
			return false
		}

		return true
	})
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] >= intervals[j][0] {
			return false
		}

		return true
	})

	ans := make([][]int, 0)

	var start, end int = intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] == start || intervals[i][0] <= end {
			end = max(end, intervals[i][1])
		} else {
			ans = append(ans, []int{start, end})
			start, end = intervals[i][0], intervals[i][1]
		}
	}

	ans = append(ans, []int{start, end})

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
