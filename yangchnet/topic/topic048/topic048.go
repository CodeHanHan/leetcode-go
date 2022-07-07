package topic048

func rotate(matrix [][]int) {
	changed := make([][]bool, len(matrix))
	for i, _ := range changed {
		changed[i] = make([]bool, len(matrix[0]))
	}

	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if !changed[i][j] {
				x_, y_ := i, j     // 要移动的点
				x, y := y_, n-1-x_ // 将要移去的位置

				tmp1 := matrix[x_][y_] // tmp 存储要移动点的值
				tmp2 := matrix[x][y]   //  tmp 存储要移去位置的值

				for !changed[x][y] {
					matrix[x][y] = tmp1  // 移动点
					changed[x][y] = true // 设置标志位
					x_, y_ = x, y        // 重新设置要移动的点
					x, y = y_, n-1-x_    // 重新计算要移去的位置
					tmp1 = tmp2
					tmp2 = matrix[x][y]
				}
			}
		}
	}
}

func rotate_1(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] =
				matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
		}
	}
}

func rotate2(matrix [][]int) [][]int {
	if len(matrix) < 1 {
		return matrix
	}

	m, n := len(matrix), len(matrix[0])
	ans := make([][]int, n)
	for i, _ := range ans {
		ans[i] = make([]int, m)
	}

	for i := m - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			ans[j][m-i-1] = matrix[i][j]
		}
	}

	return ans
}

func rotate3(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < n; i++ {
		for j, k := 0, n-1; j <= k; j, k = j+1, k-1 {
			matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
		}
	}
}
