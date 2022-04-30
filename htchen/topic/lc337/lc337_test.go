package lc337

import (
	"testing"

	tree "github.com/CodeHanHan/leetcode-go/base/Tree"
	"github.com/stretchr/testify/require"
)

func Test_rob(t *testing.T) {
	require.Equal(t, 9,
		rob(tree.BuildTree([]int{3, 4, 1, 3, 5, 1}, []int{1, 4, 3, 3, 5, 1})))
}
