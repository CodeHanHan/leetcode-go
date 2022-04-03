package goden1606

import (
	"sort"
)

func smallestDifference(a []int, b []int) int {
	sort.Ints(a)
	sort.Ints(b)

	i, j := 0, 0
	min := abs(a[i] - b[j])
	if min == 0 {
		return min
	}
	for i < len(a)-1 && j < len(b)-1 {
		if abs(a[i+1]-b[j]) > abs(a[i]-b[j+1]) {
			j++
		} else if abs(a[i+1]-b[j]) < abs(a[i]-b[j+1]) {
			i++
		} else {
			i++
			j++
		}
		if abs(a[i]-b[j]) < min {
			min = abs(a[i] - b[j])
		}
	}

	return min
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
