package topic235

// 解法1
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	a, b := p.Val, q.Val

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
		if a < pre[0] && b > pre[0] || a > pre[0] && b < pre[0] || a == pre[0] || b == pre[0] { // 在两侧, 或某一个等于root
			target = pre[0]
			return
		} else if a < pre[0] && b < pre[0] {
			index := indexOf(in, pre[0])
			fn(pre[1:index+1], in[0:index])
		} else if a > pre[0] && b > pre[0] {
			index := indexOf(in, pre[0])
			fn(pre[index+1:], in[index+1:])
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

	if root.Val < target {
		return findNode(root.Right, target)
	} else {
		return findNode(root.Left, target)
	}
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

// 解法2
func lowestCommonAncestor_1(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if p.Val < root.Val && q.Val > root.Val || p.Val > root.Val && q.Val < root.Val || p.Val == root.Val || q.Val == root.Val {
		return root
	} else if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor_1(root.Left, p, q)
	} else if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor_1(root.Right, p, q)
	}

	return nil
}
