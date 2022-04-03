package goden1616

import "math"

func subSort(array []int) []int {
	n := len(array)
	res := []int{-1, -1}

	curMin := math.MaxInt64
	for i := n - 1; i >= 0; i-- {
		if array[i] > curMin {
			res[0] = i
			continue
		}
		curMin = array[i]
	}

	curMax := math.MinInt64
	for i := 0; i < n; i++ {
		if array[i] < curMax {
			res[1] = i
			continue
		}
		curMax = array[i]
	}

	return res
}
