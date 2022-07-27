package offer22

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/LinkList"
	"github.com/stretchr/testify/require"
)

func Test_getKthFromEnd(t *testing.T) {
	require.Equal(t, BuildListWithNoHead([]int{4, 5}), getKthFromEnd(BuildListWithNoHead([]int{1, 2, 3, 4, 5}), 2))

	require.Equal(t, BuildListWithNoHead([]int{5}), getKthFromEnd(BuildListWithNoHead([]int{1, 2, 3, 4, 5}), 1))

	require.Equal(t, BuildListWithNoHead([]int{1, 2, 3, 4, 5}), getKthFromEnd(BuildListWithNoHead([]int{1, 2, 3, 4, 5}), 5))
}
