package lc463

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_islandPerimeter(t *testing.T) {
	require.Equal(t, 16,
		islandPerimeter(
			[][]int{
				[]int{0, 1, 0, 0},
				[]int{1, 1, 1, 0},
				[]int{0, 1, 0, 0},
				[]int{1, 1, 0, 0},
			}))
}

func Test_islandPerimeter1(t *testing.T) {
	require.Equal(t, 16,
		islandPerimeter1(
			[][]int{
				[]int{0, 1, 0, 0},
				[]int{1, 1, 1, 0},
				[]int{0, 1, 0, 0},
				[]int{1, 1, 0, 0},
			}))
}
