package topic92

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/LinkList"
	"github.com/stretchr/testify/require"
)

func Test_reverseBetween(t *testing.T) {
	var list *ListNode

	list = BuildListWithNoHead([]int{1, 2, 3, 4, 5})
	require.Equal(t, "1 → 4 → 3 → 2 → 5 → nil", reverseBetween(list, 2, 4).String())

	list = BuildListWithNoHead([]int{1, 2})
	require.Equal(t, "2 → 1 → nil", reverseBetween(list, 1, 2).String())

	list = BuildListWithNoHead([]int{1, 2, 3})
	require.Equal(t, "3 → 2 → 1 → nil", reverseBetween(list, 1, 3).String())

	list = BuildListWithNoHead([]int{0})
	require.Equal(t, "0 → nil", reverseBetween(list, 0, 0).String())

	list = BuildListWithNoHead([]int{3, 5})
	require.Equal(t, "5 → 3 → nil", reverseBetween(list, 1, 2).String())
}

func Test_reverseBetween1(t *testing.T) {
	var list *ListNode

	list = BuildListWithNoHead([]int{1, 2, 3, 4, 5})
	require.Equal(t, "1 → 4 → 3 → 2 → 5 → nil", reverseBetween1(list, 2, 4).String())

	list = BuildListWithNoHead([]int{1, 2})
	require.Equal(t, "2 → 1 → nil", reverseBetween1(list, 1, 2).String())

	list = BuildListWithNoHead([]int{1, 2, 3})
	require.Equal(t, "3 → 2 → 1 → nil", reverseBetween1(list, 1, 3).String())

	list = BuildListWithNoHead([]int{0})
	require.Equal(t, "0 → nil", reverseBetween1(list, 0, 0).String())

	list = BuildListWithNoHead([]int{3, 5})
	require.Equal(t, "5 → 3 → nil", reverseBetween1(list, 1, 2).String())
}
