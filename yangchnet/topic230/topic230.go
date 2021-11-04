package topic230

func kthSmallest(root *TreeNode, k int) int {
	var index int = 0
	var ret int = -1

	var fn func(root *TreeNode)
	fn = func(root *TreeNode) {
		if root == nil {
			return
		}

		fn(root.Left)
		index += 1
		if index == k {
			ret = root.Val
		}
		fn(root.Right)
	}

	fn(root)

	return ret
}
