package goden0409

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func BSTSequences(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{{}}
	}
	res := make([][]int, 0)

	queue := []*TreeNode{root}

	var fn func(queue []*TreeNode, path []int)
	fn = func(queue []*TreeNode, path []int) {
		if len(queue) == 0 {
			res = append(res, append([]int{}, path...))
			return
		}

		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			path = append(path, cur.Val)

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}

			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}

			fn(queue, path)

			queue = queue[0 : size-1]
			queue = append(queue, cur)
			path = path[:len(path)-1]
		}
	}

	fn(queue, []int{})
	return res
}
