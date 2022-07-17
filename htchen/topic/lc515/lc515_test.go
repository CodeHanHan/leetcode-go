package lc515

import (
	"testing"

	tree "github.com/CodeHanHan/leetcode-go/base/Tree"
	"github.com/stretchr/testify/require"
)

func Test_largestValues(t *testing.T) {
	require.Equal(t, []int{1, 3, 9},
		largestValues(
			tree.BuildTree([]int{1, 3, 5, 3, 2, 9}, []int{5, 3, 3, 1, 2, 9})))
}
