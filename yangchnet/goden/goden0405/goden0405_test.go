package goden0405

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isValidBST(t *testing.T) {
	tree := BuildTree([]int{1}, []int{1})
	require.Equal(t, true, isValidBST(tree))

	tree = BuildTree([]int{2, 1, 3}, []int{1, 2, 3})
	require.Equal(t, true, isValidBST(tree))

	tree = BuildTree([]int{1, 2, 4, 8, 5, 3, 7}, []int{8, 4, 2, 5, 1, 3, 7})
	require.Equal(t, false, isValidBST(tree))

	tree = BuildTree([]int{1, 1}, []int{1, 1})
	require.Equal(t, false, isValidBST(tree))
}
