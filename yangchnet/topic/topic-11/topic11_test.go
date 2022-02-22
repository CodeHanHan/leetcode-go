package topic11

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxArea(t *testing.T) {
	require.Equal(t, 49, maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))

	require.Equal(t, 1, maxArea([]int{1, 1}))

	require.Equal(t, 16, maxArea([]int{4, 3, 2, 1, 4}))

	require.Equal(t, 2, maxArea([]int{1, 2, 1}))
}

func Test_maxArea_1(t *testing.T) {
	require.Equal(t, 49, maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))

	require.Equal(t, 1, maxArea_1([]int{1, 1}))

	require.Equal(t, 16, maxArea_1([]int{4, 3, 2, 1, 4}))

	require.Equal(t, 2, maxArea_1([]int{1, 2, 1}))
}
