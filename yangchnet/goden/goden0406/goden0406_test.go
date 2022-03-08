package goden0406

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_inorderSuccessor(t *testing.T) {
	tree := BuildTree([]int{2, 1, 3}, []int{1, 2, 3})
	n := tree.Find(2)
	require.Equal(t, 2, n.Val)
	require.Equal(t, 3, inorderSuccessor(tree, n).Val)

	tree = BuildTree([]int{2, 1, 3}, []int{1, 2, 3})
	n = tree.Find(3)
	require.Equal(t, 3, n.Val)
	require.Equal(t, (*TreeNode)(nil), inorderSuccessor(tree, n))

	tree = BuildTree([]int{5, 3, 2, 1, 4, 6}, []int{1, 2, 3, 4, 5, 6})
	n = tree.Find(3)
	require.Equal(t, 3, n.Val)
	require.Equal(t, 4, inorderSuccessor(tree, n).Val)
}
