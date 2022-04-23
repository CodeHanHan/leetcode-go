package lc240

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_searchMatrix(t *testing.T) {
	require.Equal(t, true,
		searchMatrix(
			[][]int{
				[]int{1, 4, 7, 11, 15},
				[]int{2, 5, 8, 12, 19},
				[]int{3, 6, 9, 16, 22},
				[]int{10, 13, 14, 17, 24},
				[]int{18, 21, 23, 26, 30},
			}, 5))
}
