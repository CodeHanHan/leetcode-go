package lc074

func searchMatrix(matrix [][]int, target int) bool {
	l1, r1 := 0, len(matrix)-1
	row := -1
	for l1 <= r1 {
		mid := (l1 + r1) / 2
		if matrix[mid][0] == target {
			row = mid
			break
		} else if matrix[mid][0] < target {
			l1 = mid + 1
		} else {
			r1 = mid - 1
		}
	}

	//若len(matrix)为1，会出现r1 = -1,将其修正为0
	if r1 == -1 {
		r1 = 0
	}

	if row == -1 {
		row = r1
	}

	l2, r2 := 0, len(matrix[0])-1
	for l2 <= r2 {
		mid := (l2 + r2) / 2
		if matrix[row][mid] == target {
			return true
		} else if matrix[row][mid] < target {
			l2 = mid + 1
		} else {
			r2 = mid - 1
		}
	}
	return false
}
