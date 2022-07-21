package lc129

import (
	. "github.com/CodeHanHan/leetcode-go/base/Tree"
)

func sumNumbers(root *TreeNode) int {
	var dfs func(root *TreeNode, prevSum int) int
	dfs = func(root *TreeNode, prevSum int) int {
		if root == nil {
			return 0
		}
		sum := prevSum*10 + root.Val
		if root.Left == nil && root.Right == nil {
			return sum
		}
		return dfs(root.Left, sum) + dfs(root.Right, sum)
	}
	return dfs(root, 0)
}

func sumNumbers1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	nodeQueue := []*TreeNode{root}
	valueQueue := []int{root.Val}
	for len(nodeQueue) > 0 {
		node := nodeQueue[0]
        nodeQueue=nodeQueue[1:]
		value := valueQueue[0]
        valueQueue=valueQueue[1:]
		if node.Left == nil && node.Right == nil {
			sum += value
		} else {
			if node.Left != nil {
				nodeQueue = append(nodeQueue, node.Left)
				valueQueue = append(valueQueue, value*10+node.Left.Val)
			}
			if node.Right != nil {
				nodeQueue = append(nodeQueue, node.Right)
				valueQueue = append(valueQueue, value*10+node.Right.Val)
			}
		}
	}
	return sum
}
