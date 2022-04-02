package lc114

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/Tree"
	"github.com/stretchr/testify/require"
)

func Test_flatten1(t *testing.T) {
	root := BuildTree([]int{1, 2, 3, 4, 5, 6}, []int{3, 2, 4, 1, 5, 6})
	flatten1(root)
	require.Equal(t, BuildTree([]int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3, 4, 5, 6}), root)
}

func Test_flatten2(t *testing.T) {
	root := BuildTree([]int{1, 2, 3, 4, 5, 6}, []int{3, 2, 4, 1, 5, 6})
	flatten2(root)
	require.Equal(t, BuildTree([]int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3, 4, 5, 6}), root)
}
