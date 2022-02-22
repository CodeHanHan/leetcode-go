package topic235

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_lowestCommonAncestor(t *testing.T) {
	tree1 := BuildTree(
		[]int{6, 2, 0, 4, 3, 5, 8, 7, 9},
		[]int{0, 2, 3, 4, 5, 6, 7, 8, 9},
	)

	var p, q *TreeNode
	p = findNode(tree1, 2)
	q = findNode(tree1, 8)
	require.Equal(t, 6, lowestCommonAncestor(tree1, p, q).Val)

	p = findNode(tree1, 3)
	q = findNode(tree1, 5)
	require.Equal(t, 4, lowestCommonAncestor(tree1, p, q).Val)

	p = findNode(tree1, 0)
	q = findNode(tree1, 2)
	require.Equal(t, 2, lowestCommonAncestor(tree1, p, q).Val)
}

func Test_findNode(t *testing.T) {
	tree1 := BuildTree(
		[]int{6, 2, 0, 4, 3, 5, 8, 7, 9},
		[]int{0, 2, 3, 4, 5, 6, 7, 8, 9},
	)

	require.Equal(t, 6, findNode(tree1, 6).Val)

	require.Equal(t, 0, findNode(tree1, 0).Val)
}

func Test_lowestCommonAncestor_1(t *testing.T) {
	tree1 := BuildTree(
		[]int{6, 2, 0, 4, 3, 5, 8, 7, 9},
		[]int{0, 2, 3, 4, 5, 6, 7, 8, 9},
	)

	var p, q *TreeNode
	p = findNode(tree1, 2)
	q = findNode(tree1, 8)
	require.Equal(t, 6, lowestCommonAncestor_1(tree1, p, q).Val)

	p = findNode(tree1, 3)
	q = findNode(tree1, 5)
	require.Equal(t, 4, lowestCommonAncestor_1(tree1, p, q).Val)

	p = findNode(tree1, 0)
	q = findNode(tree1, 2)
	require.Equal(t, 2, lowestCommonAncestor_1(tree1, p, q).Val)
}
