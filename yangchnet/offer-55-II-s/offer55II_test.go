package offer_55_II_s

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBalance(t *testing.T) {
	root := BuildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	require.Equal(t, isBalanced(root), true)
	require.Equal(t, isBalanced1(root), true)

	root1 := BuildTree([]int{1, 2, 3, 4, 5}, []int{2, 1, 4, 5, 3})
	require.Equal(t, isBalanced(root1), false)
	require.Equal(t, isBalanced1(root1), false)
}
