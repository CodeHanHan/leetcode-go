package lc199

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
func rightSideView(root *TreeNode) []int {
	queue := make([]*TreeNode, 0)
	res := []int{}
	if root == nil {
		return res
	}
	queue = append(queue, root)
	for len(queue) != 0 {
		flag := true
		for i := len(queue); i != 0; i-- {
			pop := queue[0]
			queue = queue[1:]
			if flag {
				res = append(res, pop.Val)
				flag = false
			}
			if pop.Right != nil {
				queue = append(queue, pop.Right)
			}
			if pop.Left != nil {
				queue = append(queue, pop.Left)
			}
		}
	}
	return res
}

func rightSideView1(root *TreeNode) []int {
	res := []int{}
	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if len(res) < level {
			res = append(res, node.Val)
		}
		dfs(node.Right, level+1)
		dfs(node.Left, level+1)
	}
	dfs(root, 1)
	return res
}
