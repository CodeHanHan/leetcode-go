package topic98

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/Tree"
	"github.com/stretchr/testify/require"
)

func Test_isValidBST(t *testing.T) {
	var tree *TreeNode

	tree = BuildTree([]int{2, 1, 3}, []int{1, 2, 3})
	require.Equal(t, true, isValidBST(tree))

	tree = BuildTree([]int{1, 2, 3}, []int{1, 2, 3})
	require.Equal(t, true, isValidBST(tree))

	tree = BuildTree([]int{5, 4, 6, 3, 7}, []int{4, 5, 3, 6, 7})
	require.Equal(t, false, isValidBST(tree))
}
