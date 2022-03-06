package goden0403

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func listOfDepth(tree *TreeNode) []*ListNode {
	m := make(map[int]*ListNode)

	var fn func(root *TreeNode, deep int, op func(n *TreeNode, deep int))
	fn = func(root *TreeNode, deep int, op func(n *TreeNode, deep int)) {
		if root == nil {
			return
		}
		op(root, deep)

		fn(root.Left, deep+1, op)
		fn(root.Right, deep+1, op)
	}

	fn(tree, 0, func(n *TreeNode, deep int) {
		p, ok := m[deep]
		if !ok {
			m[deep] = &ListNode{
				Val: n.Val,
			}
			return
		}

		for p.Next != nil {
			p = p.Next
		}

		p.Next = &ListNode{
			Val: n.Val,
		}
	})

	list := make([]*ListNode, len(m))
	for i := 0; i < len(m); i++ {
		list[i] = m[i]
	}

	return list
}
