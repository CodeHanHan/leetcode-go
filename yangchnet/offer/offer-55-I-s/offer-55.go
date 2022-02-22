package offer_55_I_s_

import "math"

func maxDepth(root *TreeNode) int {
	max := 0
	var f func(*TreeNode, int)
	f = func(root *TreeNode, deep int) {
		if root == nil {
			return
		}
		if max < deep {
			max = deep
		}
		f(root.Left, deep+1)
		f(root.Right, deep+1)
	}
	f(root, 1)
	return max
}

// 递归思想
func maxDepth_1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + int(math.Max(float64(maxDepth_1(root.Left)), float64(maxDepth_1(root.Right))))
}
