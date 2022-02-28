package goden0108

func setZeroes(matrix [][]int) {
	type loc struct {
		row int
		col int
	}
	locs := make([]*loc, 0)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				locs = append(locs, &loc{i, j})
			}
		}
	}

	for _, loc := range locs {
		zero(matrix, loc.row, loc.col)
	}
}

func zero(matrix [][]int, row, col int) {
	matrix[row] = make([]int, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		matrix[i][col] = 0
	}
}
