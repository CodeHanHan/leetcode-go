package lc054

func spiralOrder(matrix [][]int) []int {
	rows, columns := len(matrix), len(matrix[0])
	if rows == 0 || columns == 0 {
		return []int{}
	}

	total := rows * columns
	visited := make([][]bool, total)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, columns)
	}

	var (
		res         = make([]int, total)
		row, column = 0, 0
		directions  = [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}
		directIndex = 0
	)
	for i := 0; i < total; i++ {
		res[i] = matrix[row][column]
		visited[row][column] = true
		nextRow, nextColumn := row+directions[directIndex][0], column+directions[directIndex][1]
		if nextRow < 0 || nextRow >= rows || nextColumn < 0 || nextColumn >= columns || visited[nextRow][nextColumn] {
			directIndex = (directIndex + 1) % 4
		}
		row += directions[directIndex][0]
		column += directions[directIndex][1]
	}
	return res
}
