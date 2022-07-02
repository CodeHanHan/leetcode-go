package topic99

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
func recoverTree(root *TreeNode) {
	var t1, t2, pre *TreeNode = nil, nil, nil

	var inOrder func(root *TreeNode)
	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}

		inOrder(root.Left)
		if pre != nil && pre.Val > root.Val {
			if t1 == nil {
				t1 = pre
			}
			t2 = root
		}

		pre = root
		inOrder(root.Right)
	}

	inOrder(root)

	t1.Val, t2.Val = t2.Val, t1.Val

	return
}
