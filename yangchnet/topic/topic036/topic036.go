package topic036

func isValidSudoku(board [][]byte) bool {
	row := make([][]int, 9)
	column := make([][]int, 9)
	block := make([][]int, 9)

	for i := 0; i < 9; i++ {
		row[i] = make([]int, 9)
		column[i] = make([]int, 9)
		block[i] = make([]int, 9)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				continue
			}

			cruNum := board[i][j] - '0' // 将byte转为int

			if row[i][cruNum-1] != 0 {
				return false
			} else {
				row[i][cruNum-1]++
			}

			if column[j][cruNum-1] != 0 {
				return false
			} else {
				column[j][cruNum-1]++
			}

			if block[j/3+(i/3)*3][cruNum-1] != 0 {
				return false
			} else {
				block[j/3+(i/3)*3][cruNum-1]++
			}
		}
	}

	return true
}
