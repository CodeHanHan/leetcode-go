package topic114

import (
	"math"

	. "github.com/CodeHanHan/leetcode-go/base/Tree"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	head := &TreeNode{Val: math.MaxInt32, Left: nil, Right: nil}
	pre := head

	list := make([]*TreeNode, 0)
	preOrder(root, func(root *TreeNode) {
		list = append(list, root)
	})

	for _, v := range list {
		pre.Right = v
		pre = pre.Right
		v.Left = nil
	}

	root = head.Right
}

func preOrder(root *TreeNode, visit func(root *TreeNode)) {
	if root == nil {
		return
	}

	visit(root)
	preOrder(root.Left, visit)
	preOrder(root.Right, visit)
}

func flatten1(root *TreeNode) {
	if root == nil {
		return
	}

	flatten1(root.Left)
	right := root.Right
	root.Right = root.Left
	root.Left = nil
	for root.Right != nil {
		root = root.Right
	}

	flatten1(right)
	root.Right = right
}
