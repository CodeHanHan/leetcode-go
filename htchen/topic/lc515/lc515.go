package lc515

import (
	"math"

	. "github.com/CodeHanHan/leetcode-go/base/Tree"
)

func largestValues(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		p := []*TreeNode{}
		maxNode := math.MinInt64
		for j := 0; j < len(q); j++ {
			node := q[j]
			if node.Val > maxNode {
				maxNode = node.Val
			}
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		res = append(res, maxNode)
		q = p
	}
	return res
}
