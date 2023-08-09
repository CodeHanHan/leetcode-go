package topic144

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/Tree"
	"github.com/stretchr/testify/require"
)

func Test_preorderTraversal(t *testing.T) {
	treeA := BuildTree([]int{3, 4, 1, 2, 5}, []int{1, 4, 2, 3, 5})
	// treeB := BuildTree([]int{4, 1}, []int{1, 4})

	require.Equal(t, []int{3, 4, 1, 2, 5}, preorderTraversal_1(treeA))
	require.Equal(t, []int{3, 4, 1, 2, 5}, preorderTraversal_2(treeA))
	require.Equal(t, []int{3, 4, 1, 2, 5}, preorderTraversal(treeA))
}
