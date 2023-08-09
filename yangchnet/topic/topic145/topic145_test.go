package topic145

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/Tree"
	"github.com/stretchr/testify/require"
)

func Test_postorderTraversal(t *testing.T) {
	treeA := BuildTree([]int{3, 4, 1, 2, 5}, []int{1, 4, 2, 3, 5})

	require.Equal(t, []int{1, 2, 4, 5, 3}, postorderTraversal(treeA))
}
