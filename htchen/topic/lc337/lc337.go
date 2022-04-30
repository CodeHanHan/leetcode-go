package lc337

import (
	. "github.com/CodeHanHan/leetcode-go/base/Tree"
)

func rob(root *TreeNode) int {
	res := dfs(root)
	return max(res[0], res[1])
}

func dfs(node *TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}

	left := dfs(node.Left)
	right := dfs(node.Right)

	isSelect := node.Val + left[0] + right[0]
	notSelect := max(left[0], left[1]) + max(right[0], right[1])

	return []int{notSelect, isSelect}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
