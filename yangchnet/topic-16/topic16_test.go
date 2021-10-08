package topic16

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_threeSumClosest(t *testing.T) {
	require.Equal(t, 82, threeSumClosest([]int{1, 2, 4, 8, 16, 32, 64, 128}, 82))

	require.Equal(t, -82, threeSumClosest([]int{-1, -2, -4, -8, -16, -32, -64, -128}, -82))

	require.Equal(t, 2, threeSumClosest([]int{-1, 2, 1, -4}, 1))

	require.Equal(t, 0, threeSumClosest([]int{0, 0, 0}, 10))

	require.Equal(t, 12, threeSumClosest([]int{1, 2, 3, 4, 5, 6, 7}, 12))

	require.Equal(t, -2, threeSumClosest([]int{-10, -2, 10, 100}, -100))
}
