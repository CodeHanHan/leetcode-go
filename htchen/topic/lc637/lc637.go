package lc637

import (
	. "github.com/CodeHanHan/leetcode-go/base/Tree"
)

func averageOfLevels(root *TreeNode) []float64 {
	res := []float64{}
	var average float64
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		p := []*TreeNode{}
		sum := 0
		for j := 0; j < len(q); j++ {
			node := q[j]
			sum += q[j].Val
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		average = float64(sum) / float64(len(q))
		res = append(res, average)
		q = p
	}
	return res
}
