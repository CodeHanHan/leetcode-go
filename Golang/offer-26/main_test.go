package main

import (
	"testing"
)

func Test_isSubStructure(t *testing.T) {
	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	root2 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}

	if !isSubStructure(root1, root2) {
		t.Error()
	}
	//fmt.Printf("%v", isSubStructure(root1, root2))

}
