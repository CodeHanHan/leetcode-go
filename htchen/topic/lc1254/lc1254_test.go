package lc1254

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_closedIsland(t *testing.T) {
	require.Equal(t, 2,
		closedIsland(
			[][]int{
				[]int{1, 1, 1, 1, 1, 1, 1, 0},
				[]int{1, 0, 0, 0, 0, 1, 1, 0},
				[]int{1, 0, 1, 0, 1, 1, 1, 0},
				[]int{1, 0, 0, 0, 0, 1, 0, 1},
				[]int{1, 1, 1, 1, 1, 1, 1, 0},
			}))
}
