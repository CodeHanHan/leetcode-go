package lc103

import (
	. "github.com/CodeHanHan/leetcode-go/base/Tree"
)

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		vals := []int{}
		p := q
		q = []*TreeNode{}
		for j := 0; j < len(p); j++ {
			node := p[j]
			vals = append(vals, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		if i%2 == 1 {
			for i, n := 0, len(vals); i < n/2; i++ {
				vals[i], vals[n-1-i] = vals[n-1-i], vals[i]
			}
		}
		res = append(res, vals)
	}
	return res
}


