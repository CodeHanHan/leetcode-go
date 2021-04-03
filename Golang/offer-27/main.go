package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
