package goden1723

func findSquare(matrix [][]int) []int {
	n := len(matrix)
	up := make([][]int, n)
	left := make([][]int, n)
	down := make([][]int, n)
	right := make([][]int, n)

	for i := range up {
		up[i] = make([]int, n)
		left[i] = make([]int, n)
		down[i] = make([]int, n)
		right[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// 如果不是黑色区域，全部设为0
			if matrix[i][j] == 1 {
				up[i][j] = 0
				left[i][j] = 0
				continue
			}

			// up
			if i <= 0 {
				up[i][j] = 1
			} else {
				up[i][j] = up[i-1][j] + 1
			}

			// left
			if j <= 0 {
				left[i][j] = 1
			} else {
				left[i][j] = left[i][j-1] + 1
			}

		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if matrix[i][j] == 1 {
				down[i][j] = 0
				right[i][j] = 0
				continue
			}

			// down
			if i >= n-1 {
				down[i][j] = 1
			} else {
				down[i][j] = down[i+1][j] + 1
			}

			// right
			if j >= n-1 {
				right[i][j] = 1
			} else {
				right[i][j] = right[i][j+1] + 1
			}
		}
	}

	ret := make([]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			size := min(down[i][j], right[i][j])
			for size > 0 {
				if up[i+size-1][j+size-1] >= size && left[i+size-1][j+size-1] >= size { // 可以围起来
					if len(ret) <= 0 {
						ret = make([]int, 3)
					}
					if size > ret[2] {
						ret[0], ret[1] = i, j
						ret[2] = size
					}
					break
				}
				size--
			}
		}
	}

	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
