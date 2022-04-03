package goden1604

func tictactoe(board []string) string {
	n := len(board)

	if n == 1 {
		return board[0]
	}

	var curStr byte
	// 检查对角线
	if curStr = board[0][0]; curStr != ' ' {
		for i := 1; i < n; i++ {
			if board[i][i] != curStr {
				break
			}

			if i == n-1 {
				return string(curStr)
			}
		}
	}

	// 检查副对角线
	if curStr = board[n-1][0]; curStr != ' ' {
		for i := 1; i < n; i++ {
			if board[n-1-i][i] != curStr {
				break
			}

			if i == n-1 {
				return string(curStr)
			}
		}
	}

	// 检查每一行
	for i := 0; i < n; i++ {
		if curStr = board[i][0]; curStr != ' ' {
			for j := 1; j < n; j++ {
				if board[i][j] != curStr {
					break
				}

				if j == n-1 {
					return string(curStr)
				}
			}
		}
	}

	// 检查每一列
	for j := 0; j < n; j++ {
		if curStr = board[0][j]; curStr != ' ' {
			for i := 1; i < n; i++ {
				if board[i][j] != curStr {
					break
				}

				if i == n-1 {
					return string(curStr)
				}
			}
		}
	}

	// 检查是否有空缺
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == ' ' {
				return "Pending"
			}
		}
	}

	return "Draw"
}
