package topic110

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/Tree"
	"github.com/stretchr/testify/require"
)

func Test_isBalanced(t *testing.T) {
	var tree *TreeNode

	tree = BuildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	require.Equal(t, true, isBalanced(tree))
}
