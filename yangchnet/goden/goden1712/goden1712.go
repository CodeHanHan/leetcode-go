package goden1712

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
func convertBiNode(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var fn func(root *TreeNode) (head, tail *TreeNode)
	fn = func(root *TreeNode) (head, tail *TreeNode) {
		if root.Left == nil && root.Right == nil {
			return root, root
		} else if root.Left == nil {
			rh, rt := fn(root.Right)
			root.Right = rh
			rt.Right = nil
			return root, rt
		} else if root.Right == nil {
			lh, lt := fn(root.Left)
			lt.Right = root
			root.Left = nil
			return lh, root
		}

		lh, lt := fn(root.Left)
		root.Left = nil
		lt.Right = root
		rh, rt := fn(root.Right)
		root.Right = rh
		rt.Right = nil

		return lh, rt
	}

	h, _ := fn(root)
	return h
}
