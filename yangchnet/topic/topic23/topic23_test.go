package topic23

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_mergeKLists(t *testing.T) {
	require.Equal(t,
		BuildListWithNoHead([]int{1, 1, 2, 3, 4, 4, 5, 6}),
		mergeKLists(
			[]*ListNode{
				BuildListWithNoHead([]int{1, 4, 5}),
				BuildListWithNoHead([]int{1, 3, 4}),
				BuildListWithNoHead([]int{2, 6}),
			},
		))

	require.Equal(t,
		BuildListWithNoHead([]int{}),
		mergeKLists(
			[]*ListNode{},
		))

	require.Equal(t,
		BuildListWithNoHead([]int{}),
		mergeKLists(
			[]*ListNode{
				BuildListWithNoHead([]int{}),
			},
		))
}
