package lc079

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	used := make([][]bool, m)
	for i := 0; i < m; i++ {
		used[i] = make([]bool, n)
	}

	var canFind func(row, col int, i int) bool
	canFind = func(row, col int, i int) bool {
		if i == len(word) {
			return true
		}
		if row < 0 || row >= m || col < 0 || col >= n {
			return false
		}
		if used[row][col] || board[row][col] != word[i] {
			return false
		}

		used[row][col] = true
		canFindNext := canFind(row+1, col, i+1) || canFind(row-1, col, i+1) || canFind(row, col-1, i+1) || canFind(row, col+1, i+1)
		if canFindNext {
			return true
		} else {
			used[row][col] = false
			return false
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] && canFind(i, j, 0) {
				return true
			}
		}
	}
	return false
}
