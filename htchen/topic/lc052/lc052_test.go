package lc052

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_totalNQueens(t *testing.T) {
	require.Equal(t, 2, totalNQueens(4))
    
	require.Equal(t, 1, totalNQueens(1))
}
