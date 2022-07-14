package lc236

import (
	. "github.com/CodeHanHan/leetcode-go/base/Tree"
)

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

func findNode(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == target {
		return root
	}

	if node := findNode(root.Left, target); node != nil {
		return node
	}

	if node := findNode(root.Right, target); node != nil {
		return node
	}

	return nil
}
