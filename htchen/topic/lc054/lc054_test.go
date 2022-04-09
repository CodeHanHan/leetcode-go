package lc054

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_spiralOrder(t *testing.T) {
	require.Equal(t, []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		spiralOrder([][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}))

	require.Equal(t, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		spiralOrder([][]int{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, []int{9, 10, 11, 12}}))
}
