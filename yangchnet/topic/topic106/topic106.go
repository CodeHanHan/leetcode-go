package topic106

import (
	. "github.com/CodeHanHan/leetcode-go/base/Tree"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) < 1 {
		return nil
	}

	rootIdx := indexOf(inorder, postorder[len(postorder)-1])
	root := &TreeNode{
		Val:   postorder[len(postorder)-1],
		Left:  buildTree(inorder[:rootIdx], postorder[:rootIdx]),
		Right: buildTree(inorder[rootIdx+1:], postorder[rootIdx:len(postorder)-1]),
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
