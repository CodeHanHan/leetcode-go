package topic107

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/Tree"
	"github.com/stretchr/testify/require"
)

func Test_levelOrderBottom(t *testing.T) {
	var tree *TreeNode

	tree = BuildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})

	require.Equal(t, [][]int{{15, 7}, {9, 20}, {3}}, levelOrderBottom(tree))
}
