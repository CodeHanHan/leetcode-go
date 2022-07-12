package lc695

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxAreaOfIsland(t *testing.T) {
	require.Equal(t, 6,
		maxAreaOfIsland(
			[][]int{
				[]int{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				[]int{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
				[]int{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
			}))

	require.Equal(t, 0,
		maxAreaOfIsland(
			[][]int{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}))
}
