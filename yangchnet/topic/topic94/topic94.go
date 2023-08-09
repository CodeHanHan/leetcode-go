package topic94

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
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	return append(inorderTraversal(root.Left), append([]int{root.Val}, inorderTraversal(root.Right)...)...)
}

// 非递归解法
func inorderTraversal_1(root *TreeNode) []int {
	ret := make([]int, 0)

	stack := []*TreeNode{root}
	cur := root.Left

	for len(stack) != 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		node := stack[len(stack)-1]
		ret = append(ret, node.Val)
		stack = stack[:len(stack)-1]

		if node.Right != nil {
			cur = node.Right
		}

	}

	return ret

}
