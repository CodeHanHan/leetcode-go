package topic73

func setZeroes(matrix [][]int) {
	zero1row, zero1column := false, false

	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
				if i == 0 {
					zero1row = true
				}
				if j == 0 {
					zero1column = true
				}
			}
		}
	}

	for i := 1; i < m; i++ {
		if matrix[i][0] == 0 {
			for j := 0; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}

	for j := 1; j < n; j++ {
		if matrix[0][j] == 0 {
			for i := 0; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}

	if zero1row {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	if zero1column {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}

	return
}
