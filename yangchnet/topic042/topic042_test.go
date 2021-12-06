package topic042

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_trap(t *testing.T) {
	require.Equal(t, 6, trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))

	require.Equal(t, 9, trap([]int{4, 2, 0, 3, 2, 5}))

	require.Equal(t, 9, trap([]int{4, 2, 0, 3, 2, 5, 2, 1, 0}))

	require.Equal(t, 7, trap([]int{4, 2, 0, 5, 2, 3, 2, 1, 0}))

	require.Equal(t, 1, trap([]int{5, 4, 1, 2}))

	require.Equal(t, 0, trap([]int{1, 7, 8}))

	require.Equal(t, 2, trap([]int{5, 3, 7, 7}))
}
