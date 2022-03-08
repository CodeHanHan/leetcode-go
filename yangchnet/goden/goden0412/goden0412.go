package goden0412

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	count := 0
	var sumfn func(root *TreeNode, s int)
	sumfn = func(root *TreeNode, s int) {
		if root == nil {
			return
		}

		s = s + root.Val
		if s == sum {
			count++
		}

		sumfn(root.Left, s)
		sumfn(root.Right, s)
	}

	var dfs func(root *TreeNode, op func(node *TreeNode))
	dfs = func(root *TreeNode, op func(node *TreeNode)) {
		if root == nil {
			return
		}

		op(root)
		dfs(root.Left, op)
		dfs(root.Right, op)
	}

	dfs(root, func(node *TreeNode) {
		sumfn(node, 0)
	})

	return count
}
