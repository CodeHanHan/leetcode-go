package topic237

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_deleteNode(t *testing.T) {
	var l, node *ListNode

	l = BuildListWithNoHead([]int{1, 2, 3, 4})
	node = l.Next.Next
	deleteNode(node)
	require.Equal(t, BuildListWithNoHead([]int{1, 2, 4}), l)

	l = BuildListWithNoHead([]int{1, 2, 3, 4, 5, 6, 7})
	node = l.Next.Next.Next.Next
	deleteNode(node)
	require.Equal(t, BuildListWithNoHead([]int{1, 2, 3, 4, 6, 7}), l)
}
