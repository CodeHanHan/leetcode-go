package topic124

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxPathSum(t *testing.T) {
	require.Equal(t, 6, maxPathSum(BuildTree([]int{1, 2, 3}, []int{2, 1, 3})))

	require.Equal(t, 42, maxPathSum(BuildTree([]int{-10, 9, 20, 15, 7}, []int{9, -10, 15, 20, 7})))

	require.Equal(t, -3, maxPathSum(BuildTree([]int{-3}, []int{-3})))

	require.Equal(t, 2, maxPathSum(BuildTree([]int{2, -1}, []int{-1, 2})))

	require.Equal(t, -1, maxPathSum(BuildTree([]int{-2, -1}, []int{-1, -2})))

}
