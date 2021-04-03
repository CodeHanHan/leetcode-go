package offer_07_m_

// 中序线索二叉树

// 未完成

type ThreadTreeNode struct {
	Val   int
	Left  *ThreadTreeNode
	Right *ThreadTreeNode
	LTag  int
	RTag  int
}

func (root *ThreadTreeNode) inThread() {
	var pre *ThreadTreeNode = nil
	if root != nil {
		pre = root.thread(pre)
		pre.Right = nil
		pre.RTag = 1
	}
}

func (root *ThreadTreeNode) thread(pre *ThreadTreeNode) *ThreadTreeNode {
	if root != nil {
		root.Left.thread(pre)
		if root.Left == nil {
			root.Left = pre
			root.LTag = 1
		}
		if pre != nil && pre.Right == nil {
			pre.Right = root
			pre.RTag = 1
		}
		root.Right.thread(root)
	}
	return pre
}

func buildTreeThread(preorder []int, inorder []int) *ThreadTreeNode {
	if len(inorder) < 1 {
		return nil
	}
	rootIndex := indexOf(inorder, preorder[0])
	root := &ThreadTreeNode{
		Val:   preorder[0],
		Left:  buildTreeThread(preorder[1:1+rootIndex], inorder[:rootIndex]),
		Right: buildTreeThread(preorder[1+rootIndex:], inorder[rootIndex+1:]),
	}
	return root
}
