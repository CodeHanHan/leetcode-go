package topic59

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_generateMatrix(t *testing.T) {
	require.Equal(t, [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}, generateMatrix(3))

	require.Equal(t, [][]int{{1}}, generateMatrix(1))

	require.Equal(t, [][]int{{1, 2}, {4, 3}}, generateMatrix(2))
}
