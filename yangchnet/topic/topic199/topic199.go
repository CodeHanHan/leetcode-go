package topic199

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
func rightSideView(root *TreeNode) []int {
	ans := make([]int, 0)

	k := 0

	var fn func(root *TreeNode, depth int)
	fn = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}

		if depth == k {
			ans = append(ans, root.Val)
			k++
		}

		fn(root.Right, depth+1)
		fn(root.Left, depth+1)
	}

	fn(root, 0)

	return ans
}
