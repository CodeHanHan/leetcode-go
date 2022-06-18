package topic56

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_merge(t *testing.T) {
	require.Equal(t, [][]int{{1, 6}, {8, 10}, {15, 18}}, merge([][]int{{1, 3}, {1, 2}, {2, 6}, {8, 10}, {15, 18}}))

	require.Equal(t, [][]int{{1, 5}}, merge([][]int{{1, 4}, {4, 5}}))

	require.Equal(t, [][]int{{1, 4}}, merge([][]int{{1, 4}, {2, 3}}))
}
