package goden0203

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	pre := node
	node.Val = node.Next.Val
	node = node.Next
	for node.Next != nil {
		node.Val = node.Next.Val
		node = node.Next
		pre = pre.Next
	}
	pre.Next = nil
}
