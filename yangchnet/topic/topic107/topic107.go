package topic107

import (
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
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	q := []*TreeNode{root}
	ans := make([][]int, 0)

	for len(q) > 0 {
		p := []*TreeNode{}
		tmp := make([]int, 0)
		for _, v := range q {
			if v.Left != nil {
				p = append(p, v.Left)
			}
			if v.Right != nil {
				p = append(p, v.Right)
			}
			tmp = append(tmp, v.Val)
		}

		q = p
		ans = append(ans, tmp)
	}

	reverse(ans)

	return ans
}

func reverse(in [][]int) {
	if len(in) < 1 {
		return
	}
	i, j := 0, len(in)-1
	for i < j {
		in[i], in[j] = in[j], in[i]
		i++
		j--
	}

	return
}
