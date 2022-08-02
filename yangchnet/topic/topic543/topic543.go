package topic543

import (
	"math"

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
func diameterOfBinaryTree(root *TreeNode) int {
	ret := math.MinInt32

	var fn func(root *TreeNode) int
	fn = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := fn(root.Left)
		right := fn(root.Right)

		ret = max(ret, left+right)

		return 1 + max(left, right)
	}

	fn(root)

	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
