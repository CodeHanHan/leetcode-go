package topic51

func solveNQueens(n int) [][]string {
	board := make([]int, n) // 下标代表行，值为列

	ans := make([][]string, 0)

	var fn func(row int)
	fn = func(row int) {
		if row == n {
			// ans = append(ans, append([]int(nil), board...))
			ans = append(ans, genBoard(board, n))
			return
		}

		for c := 0; c < n; c++ {
			if isOk(board, row, c) {
				board[row] = c // 在第row行c列放置棋子
				fn(row + 1)
			}
		}

	}

	fn(0)

	return ans
}

func isOk(board []int, row, c int) bool {
	leftUp, rightUp := c-1, c+1
	for i := row - 1; i >= 0; i-- {
		if board[i] == c { // i 行c列有棋子, 检查当前列
			return false
		}

		if leftUp >= 0 && board[i] == leftUp { // 检查左上角
			return false
		}

		if rightUp < len(board) && board[i] == rightUp { // 检查右上角
			return false
		}
		leftUp--
		rightUp++
	}

	return true
}

func genBoard(pos []int, n int) (result []string) {
	for i := 0; i < n; i++ {
		col := pos[i]
		rowStr := ""
		for j := 0; j < n; j++ {
			if j == col {
				rowStr += "Q"
			} else {
				rowStr += "."
			}
		}
		result = append(result, rowStr)
	}
	return
}
