package topic73

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_setZeroes(t *testing.T) {
	var matrix [][]int = [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	setZeroes(matrix)
	require.Equal(t, [][]int{
		{1, 0, 1},
		{0, 0, 0},
		{1, 0, 1},
	}, matrix)

	var matrix1 [][]int = [][]int{
		{0, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	setZeroes(matrix1)
	require.Equal(t, [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 1},
	}, matrix1)
}
