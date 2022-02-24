package lc153

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findMin(t *testing.T) {

	require.Equal(t, 1, findMin([]int{3, 4, 5, 1, 2}))

	require.Equal(t, 0, findMin([]int{4, 5, 6, 7, 0, 1, 2}))

	require.Equal(t, 11, findMin([]int{11, 13, 15, 17}))

	require.Equal(t, 0, findMin([]int{0}))

	require.Equal(t, 0, findMin([]int{0, 1}))
}
