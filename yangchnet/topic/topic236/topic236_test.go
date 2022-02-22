package topic236

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_lowestCommonAncestor(t *testing.T) {
	tree1 := BuildTree(
		[]int{3, 5, 6, 2, 7, 4, 1, 0, 8},
		[]int{6, 5, 7, 2, 4, 3, 0, 1, 8},
	)

	var p, q *TreeNode
	p = findNode(tree1, 2)
	q = findNode(tree1, 8)
	require.Equal(t, 3, lowestCommonAncestor(tree1, p, q).Val)

	p = findNode(tree1, 3)
	q = findNode(tree1, 5)
	require.Equal(t, 3, lowestCommonAncestor(tree1, p, q).Val)

	p = findNode(tree1, 0)
	q = findNode(tree1, 2)
	require.Equal(t, 3, lowestCommonAncestor(tree1, p, q).Val)

	tree2 := BuildTree(
		[]int{37, -34, -100, -48, -101, -54, 48, -71, -22},
		[]int{-34, -100, 37, -54, -101, -48, -71, 48, -22},
	)

	p = findNode(tree2, -71)
	q = findNode(tree2, 48)
	require.Equal(t, 48, lowestCommonAncestor(tree2, p, q).Val)

	p = findNode(tree2, -34)
	q = findNode(tree2, -100)
	require.Equal(t, -34, lowestCommonAncestor(tree2, p, q).Val)
}

func Test_findNode(t *testing.T) {
	tree1 := BuildTree(
		[]int{3, 5, 6, 2, 7, 4, 1, 0, 8},
		[]int{6, 5, 7, 2, 4, 3, 0, 1, 8},
	)

	require.Equal(t, 6, findNode(tree1, 6).Val)

	require.Equal(t, 0, findNode(tree1, 0).Val)

	require.Equal(t, 8, findNode(tree1, 8).Val)
}

func Test_lowestCommonAncestor_1(t *testing.T) {
	tree1 := BuildTree(
		[]int{3, 5, 6, 2, 7, 4, 1, 0, 8},
		[]int{6, 5, 7, 2, 4, 3, 0, 1, 8},
	)

	var p, q *TreeNode
	p = findNode(tree1, 2)
	q = findNode(tree1, 8)
	require.Equal(t, 3, lowestCommonAncestor_1(tree1, p, q).Val)

	p = findNode(tree1, 3)
	q = findNode(tree1, 5)
	require.Equal(t, 3, lowestCommonAncestor_1(tree1, p, q).Val)

	p = findNode(tree1, 0)
	q = findNode(tree1, 2)
	require.Equal(t, 3, lowestCommonAncestor_1(tree1, p, q).Val)

	tree2 := BuildTree(
		[]int{37, -34, -100, -48, -101, -54, 48, -71, -22},
		[]int{-34, -100, 37, -54, -101, -48, -71, 48, -22},
	)

	p = findNode(tree2, -71)
	q = findNode(tree2, 48)
	require.Equal(t, 48, lowestCommonAncestor_1(tree2, p, q).Val)

	p = findNode(tree2, -34)
	q = findNode(tree2, -100)
	require.Equal(t, -34, lowestCommonAncestor_1(tree2, p, q).Val)
}
