package topic102

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/Tree"
	"github.com/stretchr/testify/require"
)

func Test_levelOrder(t *testing.T) {
	var tree *TreeNode

	tree = BuildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	require.Equal(t, [][]int{
		{3},
		{9, 20},
		{15, 7},
	}, levelOrder(tree))

	tree = BuildTree([]int{}, []int{})
	require.Equal(t, [][]int{}, levelOrder(tree))
}
