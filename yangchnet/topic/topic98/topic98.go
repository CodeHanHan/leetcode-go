package topic98

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
func isValidBST(root *TreeNode) bool {

	path := []int{}
	var inOrder func(root *TreeNode)
	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}

		inOrder(root.Left)
		path = append(path, root.Val)
		inOrder(root.Right)

		return
	}

	inOrder(root)

	for i := 1; i < len(path); i++ {
		if path[i] <= path[i-1] {
			return false
		}
	}

	return true
}
