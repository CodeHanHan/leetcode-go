package topic78

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_subsets(t *testing.T) {
	require.Equal(t, [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}}, subsets([]int{1, 2, 3}))

	require.Equal(t, [][]int{{}, {0}}, subsets([]int{0}))

	require.Equal(t, [][]int{{}, {-10}}, subsets([]int{-10}))

	require.Equal(t, [][]int{{}, {-10}, {0}, {-10, 0}}, subsets([]int{-10, 0}))
}
