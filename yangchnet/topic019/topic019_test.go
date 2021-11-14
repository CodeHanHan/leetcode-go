package topic019

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removeNthFromEnd(t *testing.T) {
	require.Equal(t, BuildListWithNoHead([]int{1, 2, 3, 5}), removeNthFromEnd(BuildListWithNoHead([]int{1, 2, 3, 4, 5}), 2))

	require.Equal(t, BuildListWithNoHead([]int{}), removeNthFromEnd(BuildListWithNoHead([]int{1}), 1))

	require.Equal(t, BuildListWithNoHead([]int{1}), removeNthFromEnd(BuildListWithNoHead([]int{1, 2}), 1))

	require.Equal(t, BuildListWithNoHead([]int{2}), removeNthFromEnd(BuildListWithNoHead([]int{1, 2}), 2))
}
