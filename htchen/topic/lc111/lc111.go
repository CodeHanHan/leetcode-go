package lc111

import (
	. "github.com/CodeHanHan/leetcode-go/base/Tree"
)

func minDepth(root *TreeNode) int {
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
			if node.Left == nil && node.Right == nil {
				return res
			}
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
