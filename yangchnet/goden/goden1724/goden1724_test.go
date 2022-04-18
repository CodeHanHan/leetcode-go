package goden1724

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getMaxMatrix(t *testing.T) {
	require.Equal(t, []int{0, 1, 0, 1}, getMaxMatrix([][]int{
		{-1, 0},
		{0, -1},
	}))

	require.Equal(t, []int{0, 0, 2, 3}, getMaxMatrix([][]int{
		{9, -8, 1, 3, -2},
		{-3, 7, 6, -2, 4},
		{6, -4, -4, 8, -7},
	}))
}
