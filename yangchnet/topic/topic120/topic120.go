package topic120

import "math"

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}

	min := math.MaxInt32

	var fn func(row, col int, sum int)
	fn = func(row, col int, sum int) {
		if row == len(triangle) {
			if sum < min {
				min = sum
			}

			return
		}

		sum += triangle[row][col]

		fn(row+1, col, sum)
		fn(row+1, col+1, sum)
	}

	fn(0, 0, 0)

	return min
}

func minimumTotal1(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}

	sum := make([]int, len(triangle))
	sum[0] = triangle[0][0]

	for i := 1; i < len(triangle); i++ {
		for j := i; j >= 0; j-- {
			if j == i {
				sum[j] = sum[j-1] + triangle[i][j]
			} else if j == 0 {
				sum[j] = sum[j] + triangle[i][j]
			} else {
				sum[j] = min(sum[j]+triangle[i][j], sum[j-1]+triangle[i][j])
			}
		}
	}

	m := sum[0]
	for _, v := range sum {
		if m > v {
			m = v
		}
	}

	return m
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
