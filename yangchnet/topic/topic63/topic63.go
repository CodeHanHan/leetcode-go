package topic63

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	row, column := len(obstacleGrid), len(obstacleGrid[0])

	if obstacleGrid[row-1][column-1] == 1 {
		return 0
	}

	route := make([][]int, row)
	for i, _ := range route {
		route[i] = make([]int, column)
	}

	for i := 0; i < column; i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}
		route[0][i] = 1
	}

	for i := 0; i < row; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		route[i][0] = 1
	}

	for i := 1; i < row; i++ {
		for j := 1; j < column; j++ {
			if obstacleGrid[i][j] == 1 {
				route[i][j] = 0
				continue
			}
			route[i][j] = route[i-1][j] + route[i][j-1]
		}
	}

	return route[row-1][column-1]
}
