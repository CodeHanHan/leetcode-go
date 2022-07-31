package topic239

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxSlidingWindow(t *testing.T) {
	require.Equal(t, []int{3, 3, 5, 5, 6, 7}, maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))

	require.Equal(t, []int{1}, maxSlidingWindow([]int{1}, 1))

	require.Equal(t, []int{1, -1}, maxSlidingWindow([]int{1, -1}, 1))

	require.Equal(t, []int{3, 3, 2, 5}, maxSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3))

	require.Equal(t, []int{10, 10, 9, 2}, maxSlidingWindow([]int{9, 10, 9, -7, -4, -8, 2, -6}, 5))
}
