package topic240

func searchMatrix(matrix [][]int, target int) bool {
	i, j := 0, len(matrix[0])-1

	for i < len(matrix) && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j--
		} else if matrix[i][j] < target {
			i++
		}
	}

	return false
}
