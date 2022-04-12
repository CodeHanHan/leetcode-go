package lc130

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_solve(t *testing.T) {
	board := [][]byte{
		[]byte{'X', 'X', 'X', 'X'},
		[]byte{'X', 'O', 'O', 'X'},
		[]byte{'X', 'X', 'O', 'X'},
		[]byte{'X', 'O', 'X', 'X'}}
	solve(board)
	require.Equal(t,
		[][]byte{
			[]byte{'X', 'X', 'X', 'X'},
			[]byte{'X', 'X', 'X', 'X'},
			[]byte{'X', 'X', 'X', 'X'},
			[]byte{'X', 'O', 'X', 'X'}}, board)
}
