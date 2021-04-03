package offer_55_II_s

import "math"

func isBalanced1(root *TreeNode) bool {
	var recur func(*TreeNode) int
	recur = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := recur(root.Left)
		if left == -1 {
			return -1
		}
		right := recur(root.Right)
		if right == -1 {
			return -1
		}
		if int(math.Abs(float64(left-right))) < 2 {
			return int(math.Max(float64(left), float64(right))) + 1
		} else {
			return -1
		}
	}

	return recur(root) != -1
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isBalanced(root.Left) && isBalanced(root.Right) &&
		abs(height(root.Left)-height(root.Right)) <= 1
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(height(root.Left), height(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}
