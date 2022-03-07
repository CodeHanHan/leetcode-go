package goden0406

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	cur := false

	var n *TreeNode
	var fn func(root *TreeNode)
	fn = func(root *TreeNode) {
		if root == nil {
			return
		}

		fn(root.Left)
		if cur {
			n = root
			cur = false
			return
		}
		if root == p {
			cur = true
		}
		fn(root.Right)
	}

	fn(root)

	return n
}
