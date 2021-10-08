package topic15

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_threeSum(t *testing.T) {
	require.Equal(t, [][]int{{-1, -1, 2}, {-1, 0, 1}}, threeSum([]int{-1, 0, 1, 2, -1, -4}))

	require.Equal(t, [][]int{}, threeSum([]int{}))

	require.Equal(t, [][]int{}, threeSum([]int{-1}))

	require.Equal(t, [][]int{{0, 0, 0}}, threeSum([]int{0, 0, 0, 0}))

	require.Equal(t, [][]int{{-2, 0, 2}, {-2, 1, 1}}, threeSum([]int{-2, 0, 1, 1, 2}))
}
