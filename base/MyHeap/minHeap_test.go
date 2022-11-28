package heap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinHeap(t *testing.T) {
	h := NewMinHeap([]int{1, 7, 3, 5, 4, 2, 8}, 7)
	require.Equal(t, []int{0, 1, 4, 2, 5, 7, 3, 8}, h.vars)

	h.Insert(9)
	require.Equal(t, []int{0, 1, 2, 4, 8, 5, 7, 3, 9}, h.vars)

	require.EqualValues(t, 1, h.Pop())
	require.EqualValues(t, 2, h.Pop())

	require.Equal(t, []int{0, 3, 5, 4, 8, 9, 7}, h.vars)

	h.Sorted()
	require.Equal(t, []int{0, 9, 8, 7, 5, 4, 3}, h.vars)

}
