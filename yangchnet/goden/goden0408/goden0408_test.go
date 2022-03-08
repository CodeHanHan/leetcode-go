package goden0408

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_lowestCommonAncestor(t *testing.T) {
	tree := BuildTree([]int{5, 3, 2, 1, 4, 6}, []int{1, 2, 3, 4, 5, 6})
	n1 := tree.Find(3)
	require.Equal(t, 3, n1.Val)
	n2 := tree.Find(4)
	require.Equal(t, 4, n2.Val)
	require.Equal(t, 3, lowestCommonAncestor(tree, n1, n2).Val)
}
