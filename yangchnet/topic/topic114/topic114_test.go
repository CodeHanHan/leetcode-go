package topic114

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/Tree"
)

func Test_flatten(t *testing.T) {
	var tree *TreeNode

	tree = BuildTree([]int{1, 2, 3}, []int{2, 1, 3})
	flatten(tree)
	t.Log(tree)
}

func Test_flatten1(t *testing.T) {
	var tree *TreeNode

	tree = BuildTree([]int{1, 2, 3}, []int{2, 1, 3})
	flatten1(tree)
	t.Log(tree)
}
