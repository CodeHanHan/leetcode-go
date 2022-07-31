package lc112

import (
	"testing"

	tree "github.com/CodeHanHan/leetcode-go/base/Tree"
	"github.com/stretchr/testify/require"
)

func Test_hasPathSum(t *testing.T) {
	require.Equal(t, true,
		hasPathSum(
			tree.BuildTree(
				[]int{5, 4, 11, 7, 2, 8, 13, 4, 1},
				[]int{7, 11, 2, 4, 5, 13, 8, 4, 1}), 22))
}

func Test_hasPathSum1(t *testing.T) {
	require.Equal(t, true,
		hasPathSum1(
			tree.BuildTree(
				[]int{5, 4, 11, 7, 2, 8, 13, 4, 1},
				[]int{7, 11, 2, 4, 5, 13, 8, 4, 1}), 22))
}
