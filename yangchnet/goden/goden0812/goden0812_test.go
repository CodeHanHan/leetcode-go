package goden0812

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_solveNQueens(t *testing.T) {
	require.Equal(t, [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}}, solveNQueens(4))
}

func Test_buildStr(t *testing.T) {
	require.Equal(t, []string{".Q..", "...Q", "Q...", "..Q."}, buildStr([]int{1, 3, 0, 2}))
}
