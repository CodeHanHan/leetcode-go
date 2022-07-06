package topic113

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
func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans := make([][]int, 0)

	var fn func(root *TreeNode, path []int, sum int)
	fn = func(root *TreeNode, path []int, sum int) {
		sum += root.Val
		path = append(path, root.Val)
		if root.Left == nil && root.Right == nil {
			if sum == targetSum {
				ans = append(ans, append([]int(nil), path...))
				return
			}

			return
		}

		if root.Left != nil {
			fn(root.Left, path, sum)
		}
		if root.Right != nil {
			fn(root.Right, path, sum)
		}
	}

	fn(root, []int{}, 0)

	return ans
}
