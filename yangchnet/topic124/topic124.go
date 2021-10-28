package topic124

import "math"

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int { // 求某个树中的最大路径（从根节点开始）
		if node == nil {
			return 0
		}

		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)

		priceNewPath := node.Val + leftGain + rightGain

		maxSum = max(maxSum, priceNewPath)

		return node.Val + max(leftGain, rightGain)
	}

	maxGain(root)
	return maxSum
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
