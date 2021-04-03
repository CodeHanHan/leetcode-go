package main

import "testing"

func Test_isSymmetric(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
	}
	if !isSymmetric(root) {
		t.Error()
	}
}
