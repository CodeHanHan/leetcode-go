package goden1712

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/Tree"
	"github.com/stretchr/testify/require"
)

func Test_convertBiNode(t *testing.T) {
	require.Equal(t,
		BuildTree([]int{1, 2, 3, 4, 5, 7, 8}, []int{1, 2, 3, 4, 5, 7, 8}),
		convertBiNode(BuildTree([]int{4, 2, 1, 3, 7, 5, 8}, []int{1, 2, 3, 4, 5, 7, 8})),
	)

	require.Equal(t,
		BuildTree([]int{1, 2, 3}, []int{1, 2, 3}),
		convertBiNode(BuildTree([]int{1, 2, 3}, []int{1, 2, 3})),
	)

	require.Equal(t,
		BuildTree([]int{}, []int{}),
		convertBiNode(BuildTree([]int{}, []int{})),
	)

	require.Equal(t,
		BuildTree([]int{0, 1, 2, 3, 4, 5, 6}, []int{0, 1, 2, 3, 4, 5, 6}),
		convertBiNode(BuildTree([]int{4, 2, 1, 0, 3, 5, 6}, []int{0, 1, 2, 3, 4, 5, 6})),
	)
}
