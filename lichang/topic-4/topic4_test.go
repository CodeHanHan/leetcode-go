package topic4

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findMedianSortedArrays(t *testing.T) {
	require.Equal(t, 2.00000, findMedianSortedArrays([]int{1, 3}, []int{2}))

	require.Equal(t, 2.50000, findMedianSortedArrays([]int{1, 2}, []int{3, 4}))

	require.Equal(t, 0.00000, findMedianSortedArrays([]int{0, 0}, []int{0, 0}))

	require.Equal(t, 1.00000, findMedianSortedArrays([]int{}, []int{1}))

	require.Equal(t, 2.00000, findMedianSortedArrays([]int{2}, []int{}))
}

func Test_findMedianSortedArrays_1(t *testing.T) {
	require.Equal(t, 2.00000, findMedianSortedArrays_1([]int{1, 3}, []int{2}))

	require.Equal(t, 2.50000, findMedianSortedArrays_1([]int{1, 2}, []int{3, 4}))

	require.Equal(t, 0.00000, findMedianSortedArrays_1([]int{0, 0}, []int{0, 0}))

	require.Equal(t, 1.00000, findMedianSortedArrays_1([]int{}, []int{1}))

	require.Equal(t, 2.00000, findMedianSortedArrays_1([]int{2}, []int{}))
}
