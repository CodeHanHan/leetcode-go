package offer_55_I_s_

func maxDepth(root *TreeNode) int {
	max := 0
	var f func(*TreeNode, int)
	f = func(root *TreeNode, deep int) {
		if root == nil {
			return
		}
		if max < deep {max = deep}
		f(root.Left, deep+1)
		f(root.Right, deep+1)
	}
	f(root, 1)
	return max
}

