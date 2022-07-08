package lc617

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/Tree"
	"github.com/stretchr/testify/require"
)

func Test_mergeTrees(t *testing.T) {
	require.Equal(t,
		BuildTree([]int{3, 4, 5, 4, 5, 7}, []int{5, 4, 4, 3, 5, 7}),
		mergeTrees(
			BuildTree([]int{1, 3, 5, 2}, []int{5, 3, 1, 2}),
			BuildTree([]int{2, 1, 4, 3, 7}, []int{1, 4, 2, 3, 7}),
        ))
}
