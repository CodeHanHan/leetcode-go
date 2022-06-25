package topic79

func exist(board [][]byte, word string) bool {
	m, n, w := len(board), len(board[0]), len(word)
	var find func(i, j int, found int) bool
	find = func(i, j int, found int) bool {
		if found == w {
			return true
		}

		if i < 0 || j < 0 || i >= m || j >= n {
			return false
		}

		if board[i][j] != word[found] {
			return false
		}

		v := board[i][j]
		board[i][j] = '#'
		defer func() {
			board[i][j] = v
		}()

		return find(i-1, j, found+1) || find(i+1, j, found+1) || find(i, j-1, found+1) || find(i, j+1, found+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if find(i, j, 0) {
				return true
			}
		}
	}

	return false
}
