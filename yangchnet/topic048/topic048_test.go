package topic048

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_rotate(t *testing.T) {
	var matrix [][]int

	matrix = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
	require.Equal(t, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}, matrix)

	matrix = [][]int{{1}}
	rotate(matrix)
	require.Equal(t, [][]int{{1}}, matrix)

	matrix = [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(matrix)
	require.Equal(t, [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}}, matrix)

	matrix = [][]int{{1, 2}, {3, 4}}
	rotate(matrix)
	require.Equal(t, [][]int{{3, 1}, {4, 2}}, matrix)
}

func Test_rotate_1(t *testing.T) {
	var matrix [][]int

	matrix = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate_1(matrix)
	require.Equal(t, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}, matrix)

	matrix = [][]int{{1}}
	rotate_1(matrix)
	require.Equal(t, [][]int{{1}}, matrix)

	matrix = [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate_1(matrix)
	require.Equal(t, [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}}, matrix)

	matrix = [][]int{{1, 2}, {3, 4}}
	rotate_1(matrix)
	require.Equal(t, [][]int{{3, 1}, {4, 2}}, matrix)
}
