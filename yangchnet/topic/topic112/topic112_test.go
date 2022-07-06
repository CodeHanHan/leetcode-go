package topic112

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/Tree"
	"github.com/stretchr/testify/require"
)

func Test_hasPathSum(t *testing.T) {
	var tree *TreeNode

	tree = BuildTree([]int{1, 2, 3}, []int{2, 1, 3})
	require.Equal(t, false, hasPathSum(tree, 5))
	require.Equal(t, true, hasPathSum(tree, 3))

	tree = BuildTree([]int{}, []int{})
	require.Equal(t, false, hasPathSum(tree, 0))

	tree = BuildTree([]int{1, 2}, []int{2, 1})
	require.Equal(t, false, hasPathSum(tree, 1))
}
