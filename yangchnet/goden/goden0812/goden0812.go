package goden0812

import "bytes"

func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	var cal8queen func(board []int, row int)
	cal8queen = func(board []int, row int) {
		if row == n {
			// TODO
			res = append(res, buildStr(board))
			// res = append(res, append(board[:0:0], board...))
			return
		}

		for column := 0; column < n; column++ {
			if isValid(board, row, column, n) {
				board[row] = column
				cal8queen(board, row+1)
			}
		}
	}

	board := make([]int, n)
	cal8queen(board, 0)

	return res
}

func buildStr(board []int) []string {
	res := make([]string, 0)
	for i := 0; i < len(board); i++ {
		buf := bytes.Buffer{}
		for j := 0; j < len(board); j++ {
			if j == board[i] {
				buf.WriteString("Q")
				continue
			}
			buf.WriteString(".")
		}
		res = append(res, buf.String())
	}

	return res
}

func isValid(board []int, row, column int, n int) bool {
	leftUp, rightUp := column-1, column+1
	for k := row - 1; k >= 0; k-- {
		if k >= 0 && board[k] == column {
			return false
		}

		if leftUp >= 0 && board[k] == leftUp {
			return false
		}

		if rightUp < n && board[k] == rightUp {
			return false
		}
		leftUp--
		rightUp++
	}

	return true
}
