package lc226

import (
	. "github.com/CodeHanHan/leetcode-go/base/Tree"
)

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}
