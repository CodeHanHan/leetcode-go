package topic199

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/Tree"
	"github.com/stretchr/testify/require"
)

func Test_rightSideView(t *testing.T) {
	var tree *TreeNode

	tree = BuildTree([]int{1, 2, 5, 3, 4}, []int{2, 5, 1, 3, 4})
	require.Equal(t, []int{1, 3, 4}, rightSideView(tree))

	tree = (*TreeNode)(nil)
	require.Equal(t, []int{}, rightSideView(tree))

	tree = BuildTree([]int{1, 3}, []int{1, 3})
	require.Equal(t, []int{1, 3}, rightSideView(tree))
}
