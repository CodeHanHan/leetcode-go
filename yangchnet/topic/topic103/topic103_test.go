package topic103

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/Tree"
	"github.com/stretchr/testify/require"
)

func Test_zigzagLevelOrder(t *testing.T) {
	var tree *TreeNode

	tree = BuildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	require.Equal(t, [][]int{
		{3},
		{20, 9},
		{15, 7},
	}, zigzagLevelOrder(tree))

	tree = BuildTree([]int{1, 2, 4, 3, 5}, []int{4, 2, 1, 3, 5})
	require.Equal(t, [][]int{
		{1},
		{3, 2},
		{4, 5},
	}, zigzagLevelOrder(tree))

	tree = BuildTree([]int{0, 2, 1, 5, 1, 4, 3, 6, -1, 8}, []int{5, 1, 1, 2, 0, 6, 3, 4, -1, 8})
	require.Equal(t, [][]int{
		{0},
		{4, 2},
		{1, 3, -1},
		{8, 6, 1, 5},
	}, zigzagLevelOrder(tree))

	tree = BuildTree([]int{}, []int{})
	require.Equal(t, [][]int{}, zigzagLevelOrder(tree))

}

func Test_reverse(t *testing.T) {
	tmp := []int{5, 1, 6, 8}
	reverse(tmp)
	require.Equal(t, []int{8, 6, 1, 5}, tmp)
}
