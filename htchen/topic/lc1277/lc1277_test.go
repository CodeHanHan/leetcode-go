package lc1277

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countSquares(t *testing.T) {
	require.Equal(t, 15,
		countSquares(
			[][]int{
				[]int{0, 1, 1, 1},
				[]int{1, 1, 1, 1},
				[]int{0, 1, 1, 1},
			}))

	require.Equal(t, 7,
		countSquares(
			[][]int{
				[]int{1, 0, 1},
				[]int{1, 1, 0},
				[]int{1, 1, 0},
			}))
}
