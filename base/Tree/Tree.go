package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BuildTree 根据前序遍历序列和中序遍历序列构造一个树，遍历序列中不可含有重复值
func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) < 1 {
		return nil
	}
	rootIndex := indexOf(inorder, preorder[0])
	root := &TreeNode{
		Val:   preorder[0],
		Left:  BuildTree(preorder[1:1+rootIndex], inorder[:rootIndex]),
		Right: BuildTree(preorder[1+rootIndex:], inorder[rootIndex+1:]),
	}
	return root
}

func (root *TreeNode) InOrder(Visit func(node *TreeNode)) {
	if root == nil {
		return
	}
	root.Left.InOrder(Visit)
	Visit(root)
	root.Right.InOrder(Visit)
}

func (root *TreeNode) PreOrder(Visit func(node *TreeNode)) {
	if root == nil {
		return
	}
	Visit(root)
	root.Left.PreOrder(Visit)
	root.Right.PreOrder(Visit)
}

func (root *TreeNode) PostOrder(Visit func(node *TreeNode)) {
	if root == nil {
		return
	}
	root.Left.PostOrder(Visit)
	root.Right.PostOrder(Visit)
	Visit(root)
}

func indexOf(order []int, val int) int {
	for i, v := range order {
		if val == v {
			return i
		}
	}
	return -1
}

func (root *TreeNode) Find(v int) *TreeNode {
	if root == nil {
		return nil
	}

	var n *TreeNode

	root.InOrder(func(node *TreeNode) {
		if node.Val == v {
			n = node
		}
	})

	return n
}

type Measurer interface {
	height() int
}

func (t *TreeNode) height() int {
	if t == nil {
		return 0
	}

	return 1 + max(t.Left.height(), t.Right.height())
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

type Balancer interface {
	measurer() bool
}

func (t *TreeNode) measurer() bool {
	if t == nil {
		return true
	}

	return t.Left.measurer() && t.Right.measurer() && abs(t.Left.height()-t.Right.height()) <= 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
