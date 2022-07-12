package lc1905

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countSubIslands(t *testing.T) {
	require.Equal(t, 3,
		countSubIslands(
			[][]int{
				[]int{1, 1, 1, 0, 0},
				[]int{0, 1, 1, 1, 1},
				[]int{0, 0, 0, 0, 0},
				[]int{1, 0, 0, 0, 0},
				[]int{1, 1, 0, 1, 1}},
			[][]int{
				[]int{1, 1, 1, 0, 0},
				[]int{0, 0, 1, 1, 1},
				[]int{0, 1, 0, 0, 0},
				[]int{1, 0, 1, 1, 0},
				[]int{0, 1, 0, 1, 0}},
		))
}
