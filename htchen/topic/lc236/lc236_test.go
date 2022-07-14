package lc236

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/Tree"

	"github.com/stretchr/testify/require"
)

func Test_lowestCommonAncestor(t *testing.T) {
	tree := BuildTree(
		[]int{3, 5, 6, 2, 7, 4, 1, 0, 8},
		[]int{6, 5, 7, 2, 4, 3, 0, 1, 8},
	)

	var p *TreeNode
	var q *TreeNode
	p = findNode(tree, 5)
	q = findNode(tree, 1)
	require.Equal(t, 3, lowestCommonAncestor(tree, p, q).Val)
}
