package goden0201

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removeDuplicateNodes(t *testing.T) {
	l1 := BuildListWithNoHead([]int{1, 2, 3, 3, 2, 1})
	require.Equal(t, BuildListWithNoHead([]int{1, 2, 3}), removeDuplicateNodes(l1))

	l2 := BuildListWithNoHead([]int{1, 1, 1, 1, 2})
	require.Equal(t, BuildListWithNoHead([]int{1, 2}), removeDuplicateNodes(l2))

	l3 := BuildListWithNoHead([]int{})
	require.Equal(t, BuildListWithNoHead([]int{}), removeDuplicateNodes(l3))

	l4 := BuildListWithNoHead([]int{1})
	require.Equal(t, BuildListWithNoHead([]int{1}), removeDuplicateNodes(l4))
}
