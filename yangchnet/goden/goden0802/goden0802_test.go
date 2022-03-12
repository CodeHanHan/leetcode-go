package goden0802

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_pathWithObstacles(t *testing.T) {
	var path [][]int

	path = pathWithObstacles([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	})
	require.Equal(t, [][]int{{0, 0}, {0, 1}, {0, 2}, {1, 2}, {2, 2}}, path)

	path = pathWithObstacles([][]int{
		{0, 0, 0},
	})
	require.Equal(t, [][]int{{0, 0}, {0, 1}, {0, 2}}, path)

	path = pathWithObstacles([][]int{{1}})
	require.Equal(t, [][]int{}, path)

	path = pathWithObstacles([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	})
	require.Equal(t, [][]int{}, path)

	path = pathWithObstacles([][]int{
		{0, 0, 0},
		{1, 1, 1},
		{0, 0, 0},
	})
	require.Equal(t, [][]int{}, path)

	path = pathWithObstacles([][]int{
		{1, 0},
	})
	require.Equal(t, [][]int{}, path)
}
