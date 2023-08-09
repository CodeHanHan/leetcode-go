package topic144

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
// 递归解法
func preorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)

	var fn func(*TreeNode)
	fn = func(root *TreeNode) {
		if root == nil {
			return
		}

		ret = append(ret, root.Val)
		fn(root.Left)
		fn(root.Right)
	}

	fn(root)

	return ret
}

// 递归解法2
func preorderTraversal_2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	return append([]int{root.Val}, append(preorderTraversal_2(root.Left), preorderTraversal_2(root.Right)...)...)
}

// 非递归解法
func preorderTraversal_1(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return []int{}
	}
	stack := make([]*TreeNode, 0) // index=0为栈底
	stack = append(stack, root)

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return ret
}
