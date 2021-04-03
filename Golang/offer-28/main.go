package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return symmetry(root.Left, root.Right)
}

func symmetry(l, r *TreeNode) bool {
	/**
	  @Des: 判断给定的两个子树是否对称
	*/
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	if l.Val != r.Val {
		return false
	}
	return symmetry(l.Left, r.Right) && symmetry(l.Right, r.Left)

}
