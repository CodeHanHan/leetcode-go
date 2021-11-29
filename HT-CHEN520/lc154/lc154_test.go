package lc154

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findMin(t *testing.T) {

	require.Equal(t, 1, findMin([]int{1, 3, 5}))

	require.Equal(t, 0, findMin([]int{2, 2, 2, 0, 1}))

	require.Equal(t, 0, findMin([]int{4, 4, 4, 4, 0, 1, 2, 3}))

	require.Equal(t, 0, findMin([]int{4, 4, 4, 4, 0, 2, 2, 2}))

	require.Equal(t, 0, findMin([]int{4, 4, 0, 4, 4, 4, 4, 4}))
}
