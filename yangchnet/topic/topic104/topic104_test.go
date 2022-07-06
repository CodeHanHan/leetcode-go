package topic104

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/Tree"
	"github.com/stretchr/testify/require"
)

func Test_maxDepth(t *testing.T) {
	var tree *TreeNode

	tree = BuildTree([]int{1, 2, 3}, []int{1, 2, 3})
	require.Equal(t, 3, maxDepth(tree))
}
