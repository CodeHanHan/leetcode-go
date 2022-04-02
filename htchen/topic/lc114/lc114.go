package lc114

import (
	. "github.com/CodeHanHan/leetcode-go/base/Tree"
)

func flatten1(root *TreeNode) {
	list := preorder(root)
	for i := 1; i < len(list); i++ {
		prev, curr := list[i-1], list[i]
		prev.Left = nil
		prev.Right = curr
	}
}

func preorder(root *TreeNode) []*TreeNode {
	list := []*TreeNode{}
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		list = append(list, node)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return list
}

func flatten2(root *TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			next := curr.Left
			pre := next
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = curr.Right
			curr.Left = nil
			curr.Right = next
		}
		curr = curr.Right
	}
}
