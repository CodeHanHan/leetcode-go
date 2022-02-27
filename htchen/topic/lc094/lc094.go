package lc094

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var inorder func(n *TreeNode)

	inorder = func(n *TreeNode) {
		if n == nil {
			return
		}

		inorder(n.Left)
		res = append(res, n.Val)
		inorder(n.Right)
	}

	inorder(root)
	return res
}
