package lc112

import (
	. "github.com/CodeHanHan/leetcode-go/base/Tree"
)

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	nodeSum := []int{}
	nodeSum = append(nodeSum, root.Val)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		nowSum := nodeSum[0]
		nodeSum = nodeSum[1:]
		if node.Left == nil && node.Right == nil {
			if nowSum == targetSum {
				return true
			} else {
				continue
			}
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
			nodeSum = append(nodeSum, node.Left.Val+nowSum)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			nodeSum = append(nodeSum, node.Right.Val+nowSum)
		}
	}
	return false
}

func hasPathSum1(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	return hasPathSum1(root.Left, targetSum-root.Val) || hasPathSum1(root.Right, targetSum-root.Val)
}
