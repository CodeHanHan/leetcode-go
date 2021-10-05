package offer26

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil && B == nil {
		return true
	}
	if B == nil || A == nil {
		return false
	}
	var result bool
	if A.Val == B.Val {
		result = judge(A, B)
	}
	if !result {
		result = isSubStructure(A.Left, B)
	}
	if !result {
		result = isSubStructure(A.Right, B)
	}
	return result
}

func judge(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	if A.Val != B.Val {
		return false
	}
	return judge(A.Left, B.Left) && judge(A.Right, B.Right)

}
