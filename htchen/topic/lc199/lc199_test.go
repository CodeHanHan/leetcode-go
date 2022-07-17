package lc199

import (
	"testing"

	tree "github.com/CodeHanHan/leetcode-go/base/Tree"
	"github.com/stretchr/testify/require"
)

func Test_rightSideView(t *testing.T) {
	require.Equal(t, []int{1, 3, 4},
		rightSideView(
			tree.BuildTree([]int{1, 2, 5, 3, 4}, []int{2, 5, 1, 3, 4})))
}
func Test_rightSideView1(t *testing.T) {
	require.Equal(t, []int{1, 3, 4},
		rightSideView1(
			tree.BuildTree([]int{1, 2, 5, 3, 4}, []int{2, 5, 1, 3, 4})))
}
