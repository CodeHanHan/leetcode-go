package goden1616

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_subSort(t *testing.T) {
	require.Equal(t, []int{3, 9}, subSort([]int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19}))
	require.Equal(t, []int{-1, -1}, subSort([]int{}))
}
