package goden1708

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_bestSeqAtIndex(t *testing.T) {
	require.Equal(t, 6, bestSeqAtIndex([]int{65, 70, 56, 75, 60, 68}, []int{100, 150, 90, 190, 95, 110}))
	require.Equal(t, 1, bestSeqAtIndex([]int{1, 2, 3, 4}, []int{4, 3, 2, 1}))
	require.Equal(t, 1, bestSeqAtIndex([]int{1}, []int{4}))
	require.Equal(t, 3, bestSeqAtIndex([]int{1, 2, 2, 2, 3}, []int{4, 5, 6, 7, 8}))
}

func Test_bestSeqAtIndex2(t *testing.T) {
	require.Equal(t, 6, bestSeqAtIndex2([]int{65, 70, 56, 75, 60, 68}, []int{100, 150, 90, 190, 95, 110}))
	require.Equal(t, 1, bestSeqAtIndex2([]int{1, 2, 3, 4}, []int{4, 3, 2, 1}))
	require.Equal(t, 1, bestSeqAtIndex2([]int{1}, []int{4}))
	require.Equal(t, 4, bestSeqAtIndex2([]int{1, 2, 3, 3, 4}, []int{1, 2, 2, 3, 4}))
	require.Equal(t, 3, bestSeqAtIndex2([]int{1, 2, 2, 2, 3}, []int{4, 5, 6, 7, 8}))
}
