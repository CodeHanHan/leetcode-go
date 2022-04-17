package goden1723

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findSquare(t *testing.T) {
	require.Equal(t, []int{1, 0, 2}, findSquare([][]int{
		{1, 0, 1},
		{0, 0, 1},
		{0, 0, 1},
	}))

	require.Equal(t, []int{0, 0, 1}, findSquare([][]int{
		{0, 1, 1},
		{1, 0, 1},
		{1, 1, 0},
	}))

	require.Equal(t, []int{}, findSquare([][]int{
		{1},
	}))
}
