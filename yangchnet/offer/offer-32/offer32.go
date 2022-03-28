package offer32

import (
	. "github.com/CodeHanHan/leetcode-go/base/Tree"
)

func levelOrder(root *TreeNode) []int {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	res := make([]int, 0)
	for len(queue) > 0 {
		node := queue[0]
		res = append(res, node.Val)
		queue = queue[1:]

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}

	}

	return res
}
