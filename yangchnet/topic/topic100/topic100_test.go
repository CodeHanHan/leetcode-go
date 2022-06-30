package topic100

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/Tree"
	"github.com/stretchr/testify/require"
)

func Test_isSameTree(t *testing.T) {
	var tree *TreeNode

	tree = BuildTree([]int{1, 2}, []int{2, 1})
	require.Equal(t, true, isSameTree(tree, tree))

	tree1 := BuildTree([]int{1, 2, 3}, []int{3, 2, 1})
	require.Equal(t, false, isSameTree(tree, tree1))
}
