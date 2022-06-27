package topic86

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/LinkList"
	"github.com/stretchr/testify/require"
)

func Test_partition(t *testing.T) {
	var list *ListNode

	list = BuildListWithNoHead([]int{1, 4, 3, 2, 5, 2})
	require.Equal(t, "1 → 2 → 2 → 4 → 3 → 5 → nil", partition(list, 3).String())

	list = BuildListWithNoHead([]int{2, 1})
	require.Equal(t, "1 → 2 → nil", partition(list, 2).String())

	list = BuildListWithNoHead([]int{5})
	require.Equal(t, "5 → nil", partition(list, 5).String())

	list = BuildListWithNoHead([]int{})
	require.Equal(t, "nil", partition(list, 1).String())

	list = BuildListWithNoHead([]int{1, 2, 4, 2, 3, 1, 5})
	require.Equal(t, "1 → 2 → 2 → 1 → 4 → 3 → 5 → nil", partition(list, 3).String())

	list = BuildListWithNoHead([]int{1, 4, 3, 0, 5, 2})
	require.Equal(t, "1 → 0 → 4 → 3 → 5 → 2 → nil", partition(list, 2).String())

	list = BuildListWithNoHead([]int{1, 1})
	require.Equal(t, "1 → 1 → nil", partition(list, 0).String())

	list = BuildListWithNoHead([]int{3, 1})
	require.Equal(t, "1 → 3 → nil", partition(list, 2).String())
}
