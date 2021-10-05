package offer26

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isSubStructure(t *testing.T) {
	treeA := BuildTree([]int{3, 4, 1, 2, 5}, []int{1, 4, 2, 3, 5})
	treeB := BuildTree([]int{4, 1}, []int{1, 4})

	require.True(t, true, isSubStructure(treeA, treeB))
}

func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) < 1 {
		return nil
	}
	rootIndex := indexOf(inorder, preorder[0])
	root := &TreeNode{
		Val:   preorder[0],
		Left:  BuildTree(preorder[1:1+rootIndex], inorder[:rootIndex]),
		Right: BuildTree(preorder[1+rootIndex:], inorder[rootIndex+1:]),
	}
	return root
}

func indexOf(order []int, val int) int {
	for i, v := range order {
		if val == v {
			return i
		}
	}
	return -1
}
