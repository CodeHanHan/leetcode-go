package tree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var tree, emptyTree *TreeNode

func init() {
	tree = BuildTree([]int{3, 4, 1, 2, 5}, []int{1, 4, 2, 3, 5})
	emptyTree = BuildTree([]int{}, []int{})
}

func Test_BuildTree(t *testing.T) {
	nums_pre := make([]int, 0)
	tree.PreOrder(func(n *TreeNode) {
		nums_pre = append(nums_pre, n.Val)
	})

	nums_in := make([]int, 0)
	tree.InOrder(func(n *TreeNode) {
		nums_in = append(nums_in, n.Val)
	})

	nums_post := make([]int, 0)
	tree.PostOrder(func(n *TreeNode) {
		nums_post = append(nums_post, n.Val)
	})

	require.Equal(t, []int{3, 4, 1, 2, 5}, nums_pre)
	require.Equal(t, []int{1, 4, 2, 3, 5}, nums_in)
	require.Equal(t, []int{1, 2, 4, 5, 3}, nums_post)
}

func Test_height(t *testing.T) {
	require.Equal(t, 3, tree.height())
	require.Equal(t, 0, emptyTree.height())
}

func Test_balance(t *testing.T) {
	require.Equal(t, true, tree.measurer())
	require.Equal(t, true, emptyTree.measurer())
}
