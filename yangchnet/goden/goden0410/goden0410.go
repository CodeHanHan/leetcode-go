package goden0410

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
func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {
	if t2 == nil {
		return true
	}

	var match func(root *TreeNode, t2 *TreeNode) bool
	match = func(root *TreeNode, t2 *TreeNode) bool {
		if t2 == nil {
			return true
		}

		if root.Val != t2.Val {
			return false
		}

		return root.Val == t2.Val && match(root.Left, t2.Left) && match(root.Right, t2.Right)
	}

	var dfs func(root *TreeNode, op func(node *TreeNode))
	dfs = func(root *TreeNode, op func(node *TreeNode)) {
		if root == nil {
			return
		}

		op(root)
		dfs(root.Left, op)
		dfs(root.Right, op)
	}

	f := false
	dfs(t1, func(node *TreeNode) {
		if match(node, t2) {
			f = true
		}
	})

	return f
}
