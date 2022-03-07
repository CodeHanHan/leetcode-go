package lc078

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_subsets(t *testing.T) {
	require.Equal(t, [][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}},
		subsets([]int{1, 2, 3}))

	require.Equal(t, [][]int{{}, {0}},
		subsets([]int{0}))
}
