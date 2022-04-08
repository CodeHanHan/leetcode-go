package goden1621

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findSwapValues(t *testing.T) {
	require.Equal(t, []int{4, 6}, findSwapValues([]int{4, 1, 2, 1, 1, 2}, []int{3, 6, 3, 3}))
	require.Equal(t, []int{}, findSwapValues([]int{1, 2, 3}, []int{4, 5, 6}))
}
