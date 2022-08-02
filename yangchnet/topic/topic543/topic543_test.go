package topic543

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/Tree"
	"github.com/stretchr/testify/require"
)

func Test_diameterOfBinaryTree(t *testing.T) {
	var tree *TreeNode

	tree = BuildTree([]int{1, 2, 4, 5, 3}, []int{4, 2, 5, 1, 3})
	require.Equal(t, 3, diameterOfBinaryTree(tree))
}
