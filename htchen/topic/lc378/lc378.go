package lc378

import (
	"sort"
)

func kthSmallest1(matrix [][]int, k int) int {
	rows, columns := len(matrix), len(matrix[0])
	sorted := make([]int, rows*columns)
	i := 0
	for _, row := range matrix {
		for _, num := range row {
			sorted[i] = num
			i++
		}
	}
	sort.Ints(sorted)
	return sorted[k-1]
}

func kthSmallest2(matrix [][]int, k int) int {
	n := len(matrix)
	left, right := matrix[0][0], matrix[n-1][n-1]
	for left < right {
		mid := (left + right) / 2
		if check(matrix, mid, k, n) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func check(matrix [][]int, mid, k, n int) bool {
	i, j := n-1, 0
	cnt := 0
	for i >= 0 && j <= n-1 {
		if matrix[i][j] <= mid {
			cnt += i + 1
			j++
		} else {
			i--
		}
	}
	return cnt >= k
}
