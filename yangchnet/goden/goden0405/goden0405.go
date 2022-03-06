package goden0405

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var fn func(root *TreeNode, op func(n *TreeNode))
	fn = func(root *TreeNode, op func(n *TreeNode)) {
		if root == nil {
			return
		}

		fn(root.Left, op)
		op(root)
		fn(root.Right, op)
	}

	inOrder := make([]int, 0)
	fn(root, func(n *TreeNode) {
		inOrder = append(inOrder, n.Val)
	})

	for i := 0; i < len(inOrder)-1; i++ {
		if inOrder[i] >= inOrder[i+1] {
			return false
		}
	}

	return true
}
