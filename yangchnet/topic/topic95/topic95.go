package topic95

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
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	return helper(1, n)
}

func helper(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	allTrees := []*TreeNode{}

	for i := start; i <= end; i++ {
		leftTrees := helper(start, i-1)
		rightTrees := helper(i+1, end)

		for _, left := range leftTrees {
			for _, right := range rightTrees {
				curTree := &TreeNode{i, left, right}
				allTrees = append(allTrees, curTree)
			}
		}
	}

	return allTrees
}
