package topic103

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
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	ans := [][]int{}

	q := [2][]*TreeNode{{root}, {}}

	cur := 0
	for len(q[0]) > 0 || len(q[1]) > 0 {
		tmp := []int{}

		for len(q[cur]) > 0 {
			cn := q[cur][0]
			if cn.Left != nil {
				q[(cur+1)%2] = append(q[(cur+1)%2], cn.Left)
			}
			if cn.Right != nil {
				q[(cur+1)%2] = append(q[(cur+1)%2], cn.Right)
			}
			tmp = append(tmp, cn.Val)
			q[cur] = q[cur][1:]
		}

		if cur == 1 {
			reverse(tmp)
			// sort.Sort(sort.Reverse(sort.IntSlice(tmp)))
		}
		ans = append(ans, tmp)

		cur = (cur + 1) % 2
	}

	return ans
}

func reverse(in []int) {
	i, j := 0, len(in)-1
	for i < j {
		in[i], in[j] = in[j], in[i]
		i++
		j--
	}
}
