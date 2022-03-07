package goden0404

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isBalanced(root.Left) &&
		isBalanced(root.Right) &&
		abs(height(root.Left)-height(root.Right)) <= 1

}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(height(root.Left), height(root.Right)) + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
