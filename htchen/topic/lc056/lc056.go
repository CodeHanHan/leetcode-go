package lc056

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	left, right := intervals[0][0], intervals[0][1]
	res := make([][]int, 0)
	for i := 1; i < len(intervals); i++ {
		if right >= intervals[i][0] {
			if right < intervals[i][1] {
				right = intervals[i][1]
			}
		} else {
			res = append(res, []int{left, right})
			left = intervals[i][0]
			right = intervals[i][1]
		}
	}
	res = append(res, []int{left, right})
	return res
}
