package lc124

import (
	"math"

	. "github.com/CodeHanHan/leetcode-go/base/Tree"
)

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt64
	var maxGain func(root *TreeNode) int
	maxGain = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftGain := max(maxGain(root.Left), 0)
		rightGain := max(maxGain(root.Right), 0)

		nodePath := root.Val + leftGain + rightGain
		maxSum = max(maxSum, nodePath)

		return root.Val + max(leftGain, rightGain)
	}
	maxGain(root)
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
