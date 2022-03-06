package lc047

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_permuteUnique(t *testing.T) {
	require.Equal(t, [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}},
		permuteUnique([]int{1, 1, 2}))

	require.Equal(t, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		permuteUnique([]int{1, 2, 3}))
}
