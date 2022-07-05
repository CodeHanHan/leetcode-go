package topic109

import (
	. "github.com/CodeHanHan/leetcode-go/base/LinkList"
	. "github.com/CodeHanHan/leetcode-go/base/Tree"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	nums := []int{}
	for p := head; p != nil; p = p.Next {
		nums = append(nums, p.Val)
	}

	return sortedArrayToBST(nums)
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) < 1 {
		return nil
	}
	mid := len(nums) / 2

	root := &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}

	return root
}
