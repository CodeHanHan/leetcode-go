package lc111

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/Tree"
	"github.com/stretchr/testify/require"
)

func Test_minDepth(t *testing.T) {
	require.Equal(t, 2, minDepth(
		BuildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})))

	require.Equal(t, 5, minDepth(
		BuildTree([]int{2, 3, 4, 5, 6}, []int{2, 3, 4, 5, 6})))
}
