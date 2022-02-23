package lc015

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_threeSum(t *testing.T) {

	require.Equal(t, [][]int{{-1, -1, 2}, {-1, 0, 1}}, threeSum([]int{-1, 0, 1, 2, -1, -4}))

	require.Equal(t, [][]int{}, threeSum([]int{}))

	require.Equal(t, [][]int{}, threeSum([]int{0}))
}
