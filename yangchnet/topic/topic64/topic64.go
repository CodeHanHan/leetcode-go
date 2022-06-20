package topic64

func minPathSum(grid [][]int) int {
	row, column := len(grid), len(grid[0])
	if row <= 0 || column <= 0 {
		return 0
	}

	route := make([][]int, row)
	for i, _ := range route {
		route[i] = make([]int, column)
	}
	route[0][0] = grid[0][0]

	for i := 1; i < column; i++ {
		route[0][i] = route[0][i-1] + grid[0][i]
	}

	for j := 1; j < row; j++ {
		route[j][0] = route[j-1][0] + grid[j][0]
	}

	for i := 1; i < row; i++ {
		for j := 1; j < column; j++ {
			route[i][j] = min(route[i-1][j], route[i][j-1]) + grid[i][j]
		}
	}

	return route[row-1][column-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
