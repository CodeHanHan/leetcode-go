package offer_54_s_

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func kthLargest1(root *TreeNode, k int) int {
	order := make([]int, 0)
	inOrder(root, &order)
	return order[len(order)-k]
}

func inOrder(root *TreeNode, order *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, order)
	*order = append(*order, root.Val)
	inOrder(root.Right, order)
}

// ===================================================================

func kthLargest(root *TreeNode, k int) int {
	order := make([]int, 0)
	root.inOrder(&order)
	return order[len(order)-k]
}

func (root *TreeNode)inOrder(order *[]int){
	if root == nil {
		return
	}
	root.Left.inOrder(order)
	*order = append(*order, root.Val)
	root.Right.inOrder(order)
}

func indexOf(order []int, val int) int {
	for i, v := range order {
		if val == v {
			return i
		}
	}
	return -1
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) < 1 {return nil}
	rootIndex := indexOf(inorder, preorder[0])
	root := &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:1+rootIndex], inorder[:rootIndex]),
		Right: buildTree(preorder[1+rootIndex:], inorder[rootIndex+1:]),
	}
	return root
}

