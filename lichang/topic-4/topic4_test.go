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

func Test_kthSmallElement(t *testing.T) {
	require.Equal(t, 4, kthSmallElement([]int{1, 3, 4, 9}, []int{2, 4, 6}, 4))

	require.Equal(t, 9, kthSmallElement([]int{1, 3, 4, 9}, []int{}, 4))

	require.Equal(t, 9, kthSmallElement([]int{1, 3, 4, 9}, []int{2}, 5))

	require.Equal(t, 4, kthSmallElement([]int{1, 3, 4, 9}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 7))
}

func Test_findMedianSortedArrays_2(t *testing.T) {
	require.Equal(t, 2.00000, findMedianSortedArrays_2([]int{1, 3}, []int{2}))

	require.Equal(t, 2.50000, findMedianSortedArrays_2([]int{1, 2}, []int{3, 4}))

	require.Equal(t, 0.00000, findMedianSortedArrays_2([]int{0, 0}, []int{0, 0}))

	require.Equal(t, 1.00000, findMedianSortedArrays_2([]int{}, []int{1}))

	require.Equal(t, 2.00000, findMedianSortedArrays_2([]int{2}, []int{}))

	require.Equal(t, 1.5, findMedianSortedArrays_2([]int{1, 2}, []int{-1, 3}))
}
