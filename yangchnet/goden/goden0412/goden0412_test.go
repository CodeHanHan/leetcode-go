package goden0412

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_pathSum(t *testing.T) {
	var tree *TreeNode

	tree = BuildTree(
		[]int{5, 4, 11, 7, 2, 8, 13, 4, 5, 1},
		[]int{7, 11, 2, 4, 5, 13, 8, 5, 4, 1})
	require.Equal(t, 3, pathSum(tree, 22))

	tree = BuildTree([]int{2, 1, 3}, []int{1, 2, 3})
	require.Equal(t, 2, pathSum(tree, 3))

	tree = BuildTree([]int{-2, -3}, []int{-2, -3})
	require.Equal(t, 1, pathSum(tree, -5))

}
