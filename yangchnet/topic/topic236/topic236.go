package topic236

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pre := make([]int, 0)
	in := make([]int, 0)

	preOrder(root, func(val int) {
		pre = append(pre, val)
	})

	inOrder(root, func(val int) {
		in = append(in, val)
	})

	var target int
	var fn func(pre, in []int)
	fn = func(pre, in []int) {
		ia, ib, rootIndex := indexOf(in, p.Val), indexOf(in, q.Val), indexOf(in, pre[0])

		if ia < rootIndex && ib > rootIndex || ia > rootIndex && ib < rootIndex || ia == rootIndex || ib == rootIndex { // 在两侧, 或某一个等于root
			target = pre[0]
			return
		} else if ia < rootIndex && ib < rootIndex {
			fn(pre[1:rootIndex+1], in[0:rootIndex])
		} else if ia > rootIndex && ib > rootIndex {
			fn(pre[rootIndex+1:], in[rootIndex+1:])
		}
	}

	fn(pre, in)

	return findNode(root, target)
}

func findNode(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == target {
		return root
	}

	if node := findNode(root.Left, target); node != nil {
		return node
	}

	if node := findNode(root.Right, target); node != nil {
		return node
	}

	return nil
}

func indexOf(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if target == nums[i] {
			return i
		}
	}

	return -1
}

func preOrder(root *TreeNode, visit func(val int)) {
	if root == nil {
		return
	}

	visit(root.Val)
	preOrder(root.Left, visit)
	preOrder(root.Right, visit)
}

func inOrder(root *TreeNode, visit func(val int)) {
	if root == nil {
		return
	}

	inOrder(root.Left, visit)
	visit(root.Val)
	inOrder(root.Right, visit)
}

func lowestCommonAncestor_1(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if p == root || q == root {
		return root
	}

	left := lowestCommonAncestor_1(root.Left, p, q)
	right := lowestCommonAncestor_1(root.Right, p, q)
	if left != nil && right != nil {
		return root
	} else if left != nil {
		return left
	} else if right != nil {
		return right
	}

	return nil
}
