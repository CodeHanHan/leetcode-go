package topic129

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
func sumNumbers(root *TreeNode) int {
	ans := 0

	var fn func(root *TreeNode, val int)
	fn = func(root *TreeNode, val int) {
		if root.Left == nil && root.Right == nil {
			ans += 10*val + root.Val
			return
		}

		if root.Left != nil {
			fn(root.Left, 10*val+root.Val)
		}

		if root.Right != nil {
			fn(root.Right, 10*val+root.Val)
		}
	}

	fn(root, 0)

	return ans

}
