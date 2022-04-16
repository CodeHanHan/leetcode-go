package goden1718

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_shortestSeq(t *testing.T) {
	require.Equal(t, []int{7, 10},
		shortestSeq(
			[]int{7, 5, 9, 0, 2, 1, 3, 5, 7, 9, 1, 1, 5, 8, 8, 9, 7},
			[]int{1, 5, 9}))

	require.Equal(t, []int{}, shortestSeq([]int{1, 2, 3}, []int{4}))
	require.Equal(t, []int{0, 2}, shortestSeq([]int{1, 2, 3}, []int{1, 2, 3}))
}
