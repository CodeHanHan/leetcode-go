package lc002

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_addTwoNumbers(t *testing.T) {
	var l1, l2 *ListNode
	l1 = BuildListWithNoHead([]int{2, 4, 3})
	l2 = BuildListWithNoHead([]int{5, 6, 4})
	require.Equal(t, BuildListWithNoHead([]int{7, 0, 8}), addTwoNumbers(l1, l2))

	l1 = BuildListWithNoHead([]int{9, 9, 9, 9, 9, 9, 9})
	l2 = BuildListWithNoHead([]int{9, 9, 9, 9})
	require.Equal(t, BuildListWithNoHead([]int{8, 9, 9, 9, 0, 0, 0, 1}), addTwoNumbers(l1, l2))
}
