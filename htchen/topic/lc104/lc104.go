package lc104

import (
	. "github.com/CodeHanHan/leetcode-go/base/Tree"
)

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth1(root *TreeNode) int {
	res := 0
	if root == nil {
		return 0
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		res += 1
		p := []*TreeNode{}
		for i := 0; i < len(q); i++ {
			node := q[i]
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return res
}
