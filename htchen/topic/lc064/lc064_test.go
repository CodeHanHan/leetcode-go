package lc064

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minPathSum(t *testing.T) {
	require.Equal(t, 7, minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
	
    require.Equal(t, 12, minPathSum([][]int{{1, 2, 3}, {4, 5, 6}}))
}
