package lc052

import "strings"

var res [][]string

func totalNQueens(n int) int {
	res = [][]string{}
	board := make([][]string, n)

	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}

	backtrack(board, 0)
	return len(res)
}

func backtrack(board [][]string, row int) {
	l := len(board)
	if l == row {
		temp := make([]string, l)
		for i := 0; i < l; i++ {
			temp[i] = strings.Join(board[i], "")
		}
		res = append(res, temp)
		return
	}

	for col := 0; col < l; col++ {
		if isVaild(board, row, col) == false {
			continue
		}
		board[row][col] = "Q"
		backtrack(board, row+1)
		board[row][col] = "."
	}
}

func isVaild(board [][]string, row, col int) bool {
	n := len(board)
	for i := 0; i < row; i++ {
		if board[i][col] == "Q" {
			return false
		}
	}

	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == "Q" {
			return false
		}
	}

	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == "Q" {
			return false
		}
	}

	return true
}
