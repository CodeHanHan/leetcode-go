package topic54

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_spiralOrder(t *testing.T) {
	require.Equal(t, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}, spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))

	require.Equal(t, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}, spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))

	require.Equal(t, []int{1, 2, 6, 5}, spiralOrder([][]int{{1, 2}, {5, 6}}))

	require.Equal(t, []int{1, 2}, spiralOrder([][]int{{1, 2}}))

	require.Equal(t, []int{1}, spiralOrder([][]int{{1}}))

	require.Equal(t, []int{}, spiralOrder([][]int{}))

	require.Equal(t, []int{}, spiralOrder([][]int{{}}))
}
