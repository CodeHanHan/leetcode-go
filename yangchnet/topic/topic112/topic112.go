package topic112

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
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	ans := false

	var fn func(root *TreeNode, sum int)
	fn = func(root *TreeNode, sum int) {
		sum += root.Val
		if root.Left == nil && root.Right == nil {
			if sum == targetSum {
				ans = true
				return
			}

			return
		}

		if root.Left != nil {
			fn(root.Left, sum)
		}
		if root.Right != nil {
			fn(root.Right, sum)
		}
	}

	fn(root, 0)

	return ans
}
