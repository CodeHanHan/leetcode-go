package lc073

func setZeroes(matrix [][]int) {
	len1, len2 := len(matrix), len(matrix[0])
	row0, col0 := false, false

	for _, v1 := range matrix {
		if v1[0] == 0 {
			col0 = true
		}
	}

	for _, v2 := range matrix[0] {
		if v2 == 0 {
			row0 = true
		}
	}

	for i := 1; i < len1; i++ {
		for j := 1; j < len2; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < len1; i++ {
		for j := 1; j < len2; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if col0 {
		for _, v1 := range matrix {
			v1[0] = 0
		}
	}

	if row0 {
		for j := 0; j < len2; j++ {
			matrix[0][j] = 0
		}
	}
}
