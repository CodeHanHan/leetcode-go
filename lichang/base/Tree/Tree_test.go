package tree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_BuildTree(t *testing.T) {
	tree := BuildTree([]int{3, 4, 1, 2, 5}, []int{1, 4, 2, 3, 5})

	nums_pre := make([]int, 0)
	tree.PreOrder(func(val int) {
		nums_pre = append(nums_pre, val)
	})

	nums_in := make([]int, 0)
	tree.InOrder(func(val int) {
		nums_in = append(nums_in, val)
	})

	nums_post := make([]int, 0)
	tree.PostOrder(func(val int) {
		nums_post = append(nums_post, val)
	})

	require.Equal(t, []int{3, 4, 1, 2, 5}, nums_pre)
	require.Equal(t, []int{1, 4, 2, 3, 5}, nums_in)
	require.Equal(t, []int{1, 2, 4, 5, 3}, nums_post)
}
