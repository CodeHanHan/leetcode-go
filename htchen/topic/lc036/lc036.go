package lc036

func isValidSudoku(board [][]byte) bool {
	var row, col [9][9]int
	var area [3][3][9]int

	for i, v := range board {
		for j, v1 := range v {
			if v1 == '.' {
				continue
			}
			index := v1 - '1'
			row[i][index]++
			col[j][index]++
			area[i/3][j/3][index]++
			if row[i][index] > 1 || col[j][index] > 1 || area[i/3][j/3][index] > 1 {
				return false
			}
		}
	}
	return true
}
