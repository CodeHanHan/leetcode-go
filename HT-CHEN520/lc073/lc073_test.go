package lc073

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_setZeroes(t *testing.T) {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix)
	require.Equal(t, [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}, matrix)

	matrix1 := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(matrix1)
	require.Equal(t, [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}, matrix1)
}
