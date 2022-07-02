package topic99

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/Tree"
)

func Test_recoverTree(t *testing.T) {
	var tree *TreeNode

	tree = BuildTree([]int{1, 3, 2}, []int{3, 2, 1})
	recoverTree(tree)
	// t.Log(tree)

	tree = BuildTree([]int{1, 2}, []int{2, 1})
	recoverTree(tree)
	// t.Log(tree)

	tree = BuildTree([]int{1, 2, 3}, []int{3, 2, 1})
	recoverTree(tree)
	// t.Log(tree)
}
