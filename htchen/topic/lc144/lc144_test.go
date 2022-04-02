package lc144

import (
	"testing"

	tree "github.com/CodeHanHan/leetcode-go/base/Tree"
	"github.com/stretchr/testify/require"
)

func Test_preorderTraversal(t *testing.T) {
	require.Equal(t, []int{1, 2, 3},
		preorderTraversal1(tree.BuildTree([]int{1, 2, 3}, []int{1, 3, 2})))
}

func Test_preorderTraversa2(t *testing.T) {
	require.Equal(t, []int{1, 2, 3},
		preorderTraversal2(tree.BuildTree([]int{1, 2, 3}, []int{1, 3, 2})))
}
