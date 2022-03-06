package goden0402

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	var fn func(nums []int) *TreeNode
	fn = func(nums []int) *TreeNode {
		if len(nums) <= 0 {
			return nil
		}

		root := len(nums) / 2

		return &TreeNode{
			Val:   nums[root],
			Left:  fn(nums[0:root]),
			Right: fn(nums[root+1:]),
		}
	}

	return fn(nums)
}
