package offer27

import (
	"math"

	. "github.com/CodeHanHan/leetcode-go/base/Tree"
)

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	head := &TreeNode{
		Val:   math.MaxInt32,
		Left:  nil,
		Right: nil,
	}
	head.Left = copy(head.Left, root)
	return head.Left
}

func copy(head, root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	//head = root
	head = &TreeNode{
		Val:   root.Val,
		Left:  nil,
		Right: nil,
	}
	head.Right = copy(head.Right, root.Left)
	head.Left = copy(head.Left, root.Right)
	return head
}

func mirrorTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var fn func(root *TreeNode)
	fn = func(root *TreeNode) {
		if root == nil {
			return
		}
		root.Left, root.Right = root.Right, root.Left
		fn(root.Left)
		fn(root.Right)
	}

	fn(root)

	return root
}
