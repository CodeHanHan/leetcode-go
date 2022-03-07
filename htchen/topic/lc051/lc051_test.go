package lc051

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_solveNQueens(t *testing.T) {
	require.Equal(t, [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}}, solveNQueens(4))
}
