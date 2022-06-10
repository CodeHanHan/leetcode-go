package topic222

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/Tree"
	"github.com/stretchr/testify/require"
)

func Test_countNodes(t *testing.T) {
	var tree *TreeNode

	tree = BuildTree([]int{1, 2, 4, 5, 3, 6}, []int{4, 2, 5, 1, 6, 3})

	require.Equal(t, 6, countNodes(tree))
}
