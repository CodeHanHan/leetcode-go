package offer27

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/Tree"
	"github.com/stretchr/testify/require"
)

func Test_mirrorTree(t *testing.T) {
	var tree *TreeNode

	tree = BuildTree([]int{}, []int{})
	require.Equal(t, tree, mirrorTree(tree))

	tree = BuildTree([]int{2, 1, 3}, []int{1, 2, 3})
	require.Equal(t, BuildTree([]int{2, 3, 1}, []int{3, 2, 1}), mirrorTree(tree))
}

func Test_mirrorTree2(t *testing.T) {
	var tree *TreeNode

	tree = BuildTree([]int{}, []int{})
	require.Equal(t, tree, mirrorTree2(tree))

	tree = BuildTree([]int{2, 1, 3}, []int{1, 2, 3})
	require.Equal(t, BuildTree([]int{2, 3, 1}, []int{3, 2, 1}), mirrorTree2(tree))
}
