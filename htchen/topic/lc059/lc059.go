package lc059

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	var (
		row, column    = 0, 0
		directions     = [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}
		directionIndex = 0
	)

	for i := 1; i <= n*n; i++ {
		matrix[row][column] = i
		nextRow := row + directions[directionIndex][0]
		nextColumn := column + directions[directionIndex][1]
		if nextRow < 0 || nextRow >= n || nextColumn < 0 || nextColumn >= n || matrix[nextRow][nextColumn] > 0 {
			directionIndex = (directionIndex + 1) % 4
		}
		row += directions[directionIndex][0]
		column += directions[directionIndex][1]
	}
	return matrix
}