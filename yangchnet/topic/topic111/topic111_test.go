package topic111

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/Tree"
	"github.com/stretchr/testify/require"
)

func Test_minDepth(t *testing.T) {
	var tree *TreeNode

	tree = BuildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	require.Equal(t, 2, minDepth(tree))

	tree = BuildTree([]int{1, 2}, []int{2, 1})
	require.Equal(t, 2, minDepth(tree))
}
