package lc124

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/Tree"
	"github.com/stretchr/testify/require"
)

func Test_maxPathSum(t *testing.T) {
	require.Equal(t, 6,
		maxPathSum(
			BuildTree([]int{1, 2, 3}, []int{2, 1, 3})))

	require.Equal(t, 42,
		maxPathSum(
			BuildTree([]int{-10, 9, 20, 15, 7}, []int{9, -10, 15, 20, 7})))
}
