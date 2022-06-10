package topic154

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findMin(t *testing.T) {
	require.Equal(t, 1, findMin([]int{4, 5, 6, 1, 2, 3}))
	require.Equal(t, 1, findMin([]int{4, 5, 6, 1, 1, 2, 2, 3, 4}))
	require.Equal(t, 1, findMin([]int{1, 2}))
	require.Equal(t, 1, findMin([]int{2, 1, 2}))
	require.Equal(t, 2, findMin([]int{3, 2}))
	require.Equal(t, 1, findMin([]int{1}))
	require.Equal(t, 0, findMin([]int{0, 0, 0, 0, 0, 1, 0, 0}))
	require.Equal(t, 0, findMin([]int{1, 1, 1, 1, 1, 0, 1, 1}))
}
