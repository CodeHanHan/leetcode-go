package lc039

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_combinationSum(t *testing.T) {
	require.Equal(t, [][]int{{2, 2, 3}, {7}},
		combinationSum([]int{2, 3, 6, 7}, 7))

	require.Equal(t, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		combinationSum([]int{2, 3, 5}, 8))

	require.Equal(t, [][]int{},
		combinationSum([]int{2}, 1))
}
