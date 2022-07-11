package lc103

import (
	"testing"

	tree "github.com/CodeHanHan/leetcode-go/base/Tree"
	"github.com/stretchr/testify/require"
)

func Test_zigzagLevelOrder(t *testing.T) {
	require.Equal(t, [][]int{[]int{3}, []int{20, 9}, []int{15, 7}},
		zigzagLevelOrder(tree.BuildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})))
}
