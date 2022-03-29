package goden0813

import (
	"sort"
)

func pileBox(box [][]int) int {
	sort.Slice(box, func(i, j int) bool {
		return box[i][0] < box[j][0]
	})

	n := len(box)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = box[i][2]
	}

	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if box[j][0] < box[i][0] && box[j][1] < box[i][1] && box[j][2] < box[i][2] {
				nums[i] = max(nums[i], nums[j]+box[i][2])
			}
		}
		res = max(res, nums[i])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
